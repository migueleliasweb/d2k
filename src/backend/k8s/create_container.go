package k8s

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// func ContainerCreateHandler(params container.ContainerCreateParams) middleware.Responder {
// 	return container.NewContainerCreateCreated().WithPayload(&container.ContainerCreateCreatedBody{
// 		ID: "123",
// 	})
// }

func foo(ctx context.Context, clientset kubernetes.Interface, params container.ContainerCreateParams) {
	fmt.Println(
		clientset.CoreV1().Pods("").Create(
			params.HTTPRequest.Context(),
			&v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:   *params.Name,
					Labels: commonLabels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "foobar",
							Image:   "alpine",
							Command: []string{"/bin/sh"},
							Args:    []string{"-c", "while true; do echo hello; sleep 10;done"},
						},
					},
				},
			},
			metav1.CreateOptions{},
		),
	)
}

func containerCreateHandler(clientset kubernetes.Interface) func(params container.ContainerCreateParams) middleware.Responder {
	return func(params container.ContainerCreateParams) middleware.Responder {

		fmt.Println(params.Body.Entrypoint)

		return container.NewContainerCreateCreated().WithPayload(&container.ContainerCreateCreatedBody{
			ID: "123",
		})
	}
}
