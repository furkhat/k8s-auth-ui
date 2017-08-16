# Application overview

## Features

  * Create ServiceAccounts
  * View ServiceAccounts
  * View RoleBindings for particular ServiceAccount
  * View ClusterRoles
  * View Roles
  
![Demo](https://github.com/furkhat/k8s-auth-ui/blob/master/files/demo.gif "Demo Demo Demo")

# Run application

Run minikube with RBAC enabled
```
minikube start --extra-config=apiserver.Authorization.Mode=RBAC
```
Give permissions to the default service account used by application
```
kubectl create clusterrolebinding demo-clusterrolebinding --clusterrole=cluster-admin --serviceaccount=default:default
```
Build application
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
