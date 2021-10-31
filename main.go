package main

import (
	"net/http"
	"path/filepath"

	"github.com/migueleliasweb/d2k/src/backend/baseapi"
	"github.com/migueleliasweb/d2k/src/backend/docker141"
	"github.com/migueleliasweb/d2k/src/backend/k8s"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// GetClient configures and returns a Kubernetes client
func getClient() kubernetes.Interface {
	kubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func main() {

	api := baseapi.ConfigureBaseApi()

	// sets up handlers
	k8s.ApiConfigurator(getClient(), api)

	http.ListenAndServe("127.0.0.1:8081", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		docker141.Docker141Handler(
			rw,
			r,
		)

		docker141.FixOptionalParamsContainerCreate(
			rw,
			r,
		)

		api.Serve(nil).ServeHTTP(rw, r)
	}))
}
