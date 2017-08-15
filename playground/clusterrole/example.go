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

const testClusterRoleName = "demo-clusterrole"

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	clusterRolesClient := clientset.RbacV1beta1().ClusterRoles()

	clusterRole := &v1beta1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: testClusterRoleName,
		},
		Rules: []v1beta1.PolicyRule{
			{
				APIGroups: []string{apiv1.GroupName},
				Resources: []string{"pods"},
				Verbs:     []string{"get", "list"}, //or just ["*"]
			},
		},
	}

	fmt.Println("Creating ClusterRole...")
	result, err := clusterRolesClient.Create(clusterRole)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ClusterRole created. Name %s, Rules %s\n", result.Name, result.Rules[0])
	prompt()

	fmt.Println("List ClusterRoles...")
	list, err := clusterRolesClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Println(" - ", d.Name)
	}
	prompt()

	fmt.Printf("Deleting %s...\n", testClusterRoleName)
	deletePolicy := metav1.DeletePropagationForeground
	if err := clusterRolesClient.Delete(testClusterRoleName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted ClusterRole.")
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
