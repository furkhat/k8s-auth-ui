package main

import (
	"bufio"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var kubeConfig = filepath.Join(os.Getenv("HOME"), "/.kube/config")

const testRoleName = "demo-role"

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	rolesClient := clientset.RbacV1beta1().Roles(apiv1.NamespaceDefault)

	role := &v1beta1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name: testRoleName,
		},
		Rules: []v1beta1.PolicyRule{
			{
				APIGroups: []string{apiv1.GroupName},
				Resources: []string{"pods"},
				Verbs:     []string{"get", "list"}, //or just ["*"]
			},
		},
	}

	fmt.Println("Creating role...")
	result, err := rolesClient.Create(role)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Role created. Name %s, Rules %s\n", result.Name, result.Rules[0])
	prompt()

	fmt.Println("List Roles...")
	list, err := rolesClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Println(" - ", d.Name)
	}
	prompt()

	fmt.Printf("Deleting %s...\n", testRoleName)
	deletePolicy := metav1.DeletePropagationForeground
	if err := rolesClient.Delete(testRoleName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted Role.")
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
