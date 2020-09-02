package main

import (
	"flag"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	clientset "k8s.io/k8s-sample-controller/pkg/generated/clientset/versioned"
	informers "k8s.io/k8s-sample-controller/pkg/generated/informers/externalversions"
	"k8s.io/k8s-sample-controller/pkg/signals"
)

var (
	masterURL  string
	kubeConfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// set up signals so we handler the first shutdown signal gracefully
	stopCH := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeConfig)
	if err != nil {
		klog.Fatalf("error building kubeConfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientSet: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientSet: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	controller := NewController(kubeClient, exampleClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		exampleInformerFactory.Samplecontroller().V1alpha1().Foos())

	// notice that this is no need to run Start methods in a separate goroutine.(i.e. go kubeInformerFactory.Start(stopCH))
	// Start method is non-blocking and runs all registered informers a dedicated goroutine.
	kubeInformerFactory.Start(stopCH)
	exampleInformerFactory.Start(stopCH)

	if err = controller.Run(2, stopCH); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeConfig, "kubeConfig", "", "Path to a kubeConfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeConfig. Only required if out-of-cluster.")
}
