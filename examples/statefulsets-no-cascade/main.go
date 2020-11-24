package main

import (
	"flag"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfig string
	namespace  string
	name       string
)

func init() {
	flag.StringVar(&kubeConfig, "kubeConfig", "", "Give the namespace name (optional)")
	flag.StringVar(&namespace, "namespace", "", "*Give the namespace name")
	flag.StringVar(&name, "name", "", "*Give the resource name")
}

func loadKubeConfig() (*rest.Config, error) {
	if kubeConfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func main() {
	flag.Parse()

	if namespace == "" {
		log.Println("Usage")
		flag.PrintDefaults()
		log.Fatal("Namespce is not given")
	}

	if name == "" {
		log.Println("Usage")
		flag.PrintDefaults()
		log.Fatal("Resoutce name is not given")
	}

	log.Printf("Loading kube config")
	config, err := loadKubeConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	policy := metav1.DeletePropagationOrphan
	noCascade := metav1.DeleteOptions{
		PropagationPolicy: &policy,
	}
	//Note: with client-go v0.18 + need the context parameter
	err = clientset.AppsV1().StatefulSets(namespace).Delete(name, &noCascade)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted resource %s sucessfully, however dependent resources not \n", name)
}
