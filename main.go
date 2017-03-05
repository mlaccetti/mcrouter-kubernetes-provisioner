package main

import (
	"flag"
	"fmt"
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
	"k8s.io/client-go/1.5/tools/cache"
	"k8s.io/client-go/1.5/tools/clientcmd"
)

var clientset *kubernetes.Clientset = nil
var namespace *string = nil
var mcrouterConfig *string = nil

func getMemcachedPods() (map[string]string, error) {
	kubeLabelSelector, err := labels.Parse("app in (memcached)")

	if err != nil {
		return nil, err
	}

	pods, err := clientset.Core().Pods(*namespace).List(api.ListOptions{LabelSelector: kubeLabelSelector})
	if err != nil {
		return nil, err
	}

	memcachedPods := make(map[string]string)
	for _, pod := range pods.Items {
		memcachedPods[pod.Name] = pod.Status.PodIP
	}

	return memcachedPods, nil
}

func updateTemplate(pods map[string]string) (error) {
	err := lib.Parse("./template/mcrouter-config.tpl", *mcrouterConfig, pods)
	if err != nil {
		return err
	}

	return nil
}

func podsModified(obj interface{}) {
	pod := obj.(*v1.Pod)

	if pod.ObjectMeta.Labels["app"] == "memcached" {
		pods, err := getMemcachedPods()
		if err != nil {
			fmt.Println("We could not get a list of memcached pods. ", err)
		}

		fmt.Println("memcached pods: ", pods)

		err = updateTemplate(pods)
		if err != nil {
			fmt.Println("We could not update mcrouter config. ", err)
		}
	}
}

func watchPods() {
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
	// uses the current context in kubeconfig
	usr, _ := user.Current()
	dir := usr.HomeDir
	file := strings.Join([]string{dir, "/.kube/config"}, "")

	// load our flags
	namespace = flag.String("namespace", "", "namespace in kubernetes to find memcached pods")
	kubeconfig := flag.String("kubeconfig", file, "absolute path to the kubeconfig file")
	mcrouterConfig = flag.String("mcrouterConfig", "mcrouter-config.json", "absolute path to the mcrouter config json location")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watchPods()

	runtime.Goexit()

	fmt.Println("Exit")
}
