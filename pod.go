package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Create a new Kubernetes config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/saiyam/Downloads/civo-cost-calculator-kubeconfig")
	if err != nil {
		panic(err)
	}

	// Create a new Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Define the pod
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-pod",
			Labels: map[string]string{
				"app": "my-app",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "my-container",
					Image: "nginx:latest",
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}

	// Create the pod
	ctx := context.TODO()
	result, err := clientset.CoreV1().Pods("default").Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created pod %q.\n", result.GetObjectMeta().GetName())
}

