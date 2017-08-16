package application

import (
	"html/template"
	"log"
	"path/filepath"

	"github.com/furkhat/k8s-users/webapp/k8s_clients"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Application struct {
	ServiceAccountsClient k8s_clients.ServiceAccountsClientInterface
	RolesClient           k8s_clients.RolesClientInterface
	ClusterRolesClient    k8s_clients.ClusterRolesClientInterface
	NamespacesClient      k8s_clients.NamespacesClientInterface
	RoleBindingsClient    k8s_clients.RoleBindingsClientInterface
	TemplateBuilder       *TemplatesBuilder
}

func NewApplication(appConfig *ApplicationConfig) *Application {
	config, err := clientcmd.BuildConfigFromFlags("", appConfig.KubeConfigPath)
	if err != nil {
		log.Panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}

	return &Application{
		ServiceAccountsClient: k8s_clients.NewServiceAccountsClient(clientset),
		RolesClient:           k8s_clients.NewRolesClient(clientset),
		ClusterRolesClient:    k8s_clients.NewClusterRolesClient(clientset),
		NamespacesClient:      k8s_clients.NewNamespacesClient(clientset),
		RoleBindingsClient:    k8s_clients.NewRoleBindingsClient(clientset),
		TemplateBuilder:       &TemplatesBuilder{appConfig.TemplatesDir},
	}
}

type TemplatesBuilder struct {
	templateDir string
}

func (builder *TemplatesBuilder) Build(templateName string) *template.Template {
	return template.Must(
		template.ParseFiles(
			filepath.Join(builder.templateDir, "base.html"),
			filepath.Join(builder.templateDir, templateName+".html"),
		),
	)
}
