package main

import (
	"flag"
	"fmt"
	"os/user"
	"strings"
	"time"

	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/pkg/fields"
	"k8s.io/client-go/1.5/pkg/util/wait"
	"k8s.io/client-go/1.5/tools/cache"
	"k8s.io/client-go/1.5/tools/clientcmd"
	"runtime"
)

func podCreated(obj interface{}) {
	pod := obj.(*v1.Pod)
	fmt.Println("Pod created: " + pod.ObjectMeta.Name)
	fmt.Println("Pod labels: ", pod.ObjectMeta.Labels)
}

func podDeleted(obj interface{}) {
	pod := obj.(*v1.Pod)
	fmt.Println("Pod deleted: " + pod.ObjectMeta.Name)
}

func watchPods(client *kubernetes.Clientset) {
	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(client.CoreClient, "pods", api.NamespaceAll, fields.Everything())

	resyncPeriod := 30 * time.Minute

	//Setup an informer to call functions when the watchlist changes
	_, eController := cache.NewInformer(
		watchlist,
		&v1.Pod{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    podCreated,
			DeleteFunc: podDeleted,
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

	kubeconfig := flag.String("kubeconfig", file, "absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watchPods(clientset)

	runtime.Goexit()

	fmt.Println("Exit")
}
