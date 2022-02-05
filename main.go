package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	var conda_env *string

	home := homedir.HomeDir()

	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[path to the kubeconfig file]")
		conda_env = flag.String("environment", "", "[path to your conda environment.yaml]")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "[path to the kubeconfig file]")
		conda_env = flag.String("environment", "", "[path to your conda environment.yaml]")
	}

	flag.Parse()

	env_file, err := ioutil.ReadFile(*conda_env)
	if err != nil {
		panic("Error reading environment file")
	}

	//print out the content for now
	//TODO: process this enviroment file
	fmt.Println(string(env_file))

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	if clientset != nil {
		fmt.Println("just a placeholder print statment")
	}
}
