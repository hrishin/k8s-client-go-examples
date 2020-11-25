# Delete a resource object with cascade=false option

This is an example to demonstrate how to delete the resource but with deleting its child resources.
Its ~ of doing `kubectl delete <kind/resource> --cascade=true`.
Intent is to set the `[PropagationPolicy](https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1#DeleteOptions)` delete option.

## Build
```
➜ statefulsets-no-cascade git:(mock-watch) ✗ make build    
go build -o delete-stateful main.go
```

## Test

Let's create a Statefulset

```
➜ statefulsets-no-cascade git:(mock-watch) ✗ kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/staging/volumes/vsphere/simple-statefulset.yaml
service/nginx created
statefulset.apps/web created
```
Which has created the the `Statefulset` and its `Pod` resources

```
➜ statefulsets-no-cascade git:(mock-watch) ✗ kubectl get all -l "app=nginx"         
NAME    READY  STATUS  RESTARTS  AGE
pod/web-0  0/1   Pending  0     5m2s

NAME      TYPE    CLUSTER-IP  EXTERNAL-IP  PORT(S)  AGE
service/nginx  ClusterIP  None     <none>    80/TCP  5m2s

NAME          READY  AGE
statefulset.apps/web  0/14  5m2s
```


Now let's delete the pod however retains it's associated resources.
```

➜ statefulsets-no-cascade git:(mock-watch) ✗ ./delete-stateful --namespace default --name=web --kubeConfig=$HOME/.kube/config
2020/11/25 21:04:53 Loading kube config
2020/11/25 21:04:53 Deleted the resource <web> sucessfully, however dependent resources are not deleted such as pods
```

Now check its resources.
```
➜ statefulsets-no-cascade git:(mock-watch) ✗ kubectl get all -l "app=nginx"                          
NAME    READY  STATUS  RESTARTS  AGE
pod/web-0  0/1   Pending  0     7m38s

NAME      TYPE    CLUSTER-IP  EXTERNAL-IP  PORT(S)  AGE
service/nginx  ClusterIP  None     <none>    80/TCP  7m38s
```