package main

import (
	"flag"
	logger "log"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/mlaccetti/mcrouter-kubernetes-provisioner/lib"
	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/pkg/fields"
	"k8s.io/client-go/1.5/pkg/labels"
	"k8s.io/client-go/1.5/pkg/util/wait"
	"k8s.io/client-go/1.5/rest"
	"k8s.io/client-go/1.5/tools/cache"
	"k8s.io/client-go/1.5/tools/clientcmd"
)

var log = logger.New(os.Stdout, "[mkp]", logger.Ldate | logger.Ltime | logger.Lmicroseconds | logger.Llongfile)

var clientset *kubernetes.Clientset = nil
var namespace *string = nil
var mcrouterConfig *string = nil

func getMemcachedPods() (map[string]string, error) {
	log.Println("Getting memcached pods from kubernetes.")

	kubeLabelSelector, err := labels.Parse("app in (memcached)")

	if err != nil {
		log.Fatal("Could not create kubernetes label selector.", err)
		return nil, err
	}

	pods, err := clientset.Core().Pods(*namespace).List(api.ListOptions{LabelSelector: kubeLabelSelector})
	if err != nil {
		log.Fatal("Could not get memcached pods from kubernetes.", err)
		return nil, err
	}

	memcachedPods := make(map[string]string)
	for _, pod := range pods.Items {
		memcachedPods[pod.Name] = pod.Status.PodIP
	}

	log.Println("Retrieved memcached pods from kubernetes.")
	return memcachedPods, nil
}

func updateConfigFile(pods map[string]string) (error) {
	log.Println("Updating config file.")

	err := lib.Parse("./template/mcrouter-config.tpl", *mcrouterConfig, pods)
	if err != nil {
		log.Fatal("Could not create load mcrouter config template.", err)
		return err
	}

	log.Println("Config file updated.")
	return nil
}

func podsModified(obj interface{}) {
	pod := obj.(*v1.Pod)

	if pod.ObjectMeta.Labels["app"] == "memcached" {
		log.Println("Detected pod modification.")

		pods, err := getMemcachedPods()
		if err != nil {
			log.Println("We could not get a list of memcached pods. ", err)
		}

		log.Println("memcached pods: ", pods)

		err = updateConfigFile(pods)
		if err != nil {
			log.Fatal("We could not update mcrouter config. ", err)
		}
	}
}

func watchPods() {
	log.Println("Configuring pod watcher.")

	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(clientset.CoreClient, "pods", api.NamespaceAll, fields.Everything())

	resyncPeriod := 30 * time.Minute

	//Setup an informer to call functions when the watchlist changes
	_, eController := cache.NewInformer(
		watchlist,
		&v1.Pod{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    podsModified,
			DeleteFunc: podsModified,
		},
	)

	//Run the controller as a goroutine
	go eController.Run(wait.NeverStop)
}

func main() {
	log.Println("Getting flags...")

	// load our flags
	inCluster := flag.Bool("incluster", true, "tell us if we are running within kubernetes or not (defaults to true)")
	namespace = flag.String("namespace", "", "namespace in kubernetes to find memcached pods")
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	mcrouterConfig = flag.String("mcrouterconfig", "mcrouter-config.json", "absolute path to the mcrouter config json location")

	flag.Parse()

	log.Printf("Processed flags: inCluster [ %b ] | namespace [ %s ] | kubeconfig [ %s ] | mcrouter config [ %s ]\n", *inCluster, *namespace, *kubeconfig, *mcrouterConfig)

	if *inCluster {
		log.Println("Running in-cluster mode, loading config.")

		// in-context
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Panic("Could not load cluster config.", err)
		}

		// creates the clientset
		log.Println("Creating in-cluster clientset.")
		clientset, err = kubernetes.NewForConfig(config)

		if err != nil {
			log.Panic("Could not create clientset.", err)
		}

		log.Println("In-cluster clientset created.")
	} else {
		log.Println("Running in external mode, checking if user passed in a kubeconfig to use.")
		if *kubeconfig == "" {
			// uses the current context in kubeconfig
			usr, _ := user.Current()
			dir := usr.HomeDir
			*kubeconfig = strings.Join([]string{dir, "/.kube/config"}, "")
		}

		log.Printf("Using %s as the kubeconfig.", *kubeconfig)

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			log.Panic("Could not load config.", err)
		}

		// creates the clientset
		log.Println("Creating external clientset.")
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Panic("Could not create external clientset.", err)
		}

		log.Println("External clientset created.")
	}

	watchPods()

	runtime.Goexit()

	log.Println("Terminating provisioner.")
}
