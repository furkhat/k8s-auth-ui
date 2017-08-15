package main

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var kubeConfig = os.Getenv("HOME") + "/.kube/config"
var testServiceAccountName = "demo-serviceaccount"

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	serviceAccountsClient := clientset.CoreV1().ServiceAccounts(apiv1.NamespaceDefault)

	serviceAccount := &apiv1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: testServiceAccountName,
		},
	}

	fmt.Println("Creating serviceaccount...")
	result, err := serviceAccountsClient.Create(serviceAccount)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created service account ", result.Name)

	fmt.Println("List serviceaccounts...")
	list, err := serviceAccountsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Println(" - ", d.Name)
	}

	fmt.Println("Deleting serviceaccount...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := serviceAccountsClient.Delete(testServiceAccountName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted serviceaccount.")
}
