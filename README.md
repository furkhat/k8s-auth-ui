# Application overview

## Features

  * Create ServiceAccounts
  * View ServiceAccounts
  * View RoleBindings for particular ServiceAccount
  * View ClusterRoles
  * View Roles
  
![Demo](https://github.com/furkhat/k8s-auth-ui/blob/master/files/demo.gif "features features features")
  
## Source Code
```./
├── files
├── playground
│   ├── clusterrole/
│   ├── role/
│   └── serviceaccount/
└── webapp
    ├── application/
    ├── handlers/
    ├── k8s_clients/
    └── templates/
    └── main.py
    
```
The [`playground/`](https://github.com/furkhat/k8s-auth-ui/tree/master/playground) folder contains example programs that demonstrate the fundamental operations for managing on resources, such as Create, List, Delete. 

There are example programs for [ClusterRoles](https://github.com/furkhat/k8s-auth-ui/blob/master/playground/clusterrole/example.go), [Roles](https://github.com/furkhat/k8s-auth-ui/blob/master/playground/role/example.go) and [ServiceAccounts](https://github.com/furkhat/k8s-auth-ui/blob/master/playground/serviceaccount/example.go)

The [`webapp/`](https://github.com/furkhat/k8s-auth-ui/tree/master/webapp)
 - `application/` application objects and app configuration
 - `handlers/` request handlers
 - `k8s_clients/` a thin abstraction layer above api clients to encapsulate and make interaction with api easier
 - `templates/` html templates
 - `main.py` entry point for the app. Creates application objects and initializes routes

# Run application

Run minikube with RBAC enabled
```
minikube start --extra-config=apiserver.Authorization.Mode=RBAC
```
Give permissions to the default service account used by application
```
kubectl create clusterrolebinding demo-clusterrolebinding --clusterrole=cluster-admin --serviceaccount=default:default
```
Clone this repo and build the app image
```
docker build -t in-cluster:v1 .
```
Run application
```
kubectl run demo --image=in-cluster:v1 --port=8080
```
Expose the pod as a kubernetes service
```
kubectl expose deployment demo --type=LoadBalancer
```
And finally type
```
minikube service demo
```
this automatically opens up a browser window using a local IP address that serves the app.
