package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rodaine/table"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	namespace := flag.String("ns", "default", "namespace of the secret") 
	secret := flag.String("s", "", "secret to fetch")
	flag.Parse()
	if len(*secret)  == 0 {
		fmt.Println("Please specify a secret to look up with -s")
		os.Exit(1)
	}
	fetchSecret(*namespace, *secret)
}


func fetchSecret(ns string, name string) {

	client := k8sconnect()
	secrets := client.CoreV1().Secrets(ns)
	secret, err := secrets.Get(context.TODO(), name  , metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	tbl := table.New("Secret", "Value")
	for name, val := range secret.Data {
		tbl.AddRow(name, string(val))
	}
	tbl.Print()
}

func k8sconnect() (*kubernetes.Clientset) {

	defaultCfg := filepath.Join(homedir.HomeDir(), ".kube", "config")

	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", defaultCfg, "kubeconfig path")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	conn, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

