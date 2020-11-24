# Delete a statefulset object with cascade=false option

## Build
```
make build
```

## Test
```
./delete-stateful --namespace=mysql --name=web --kubeConfig=$HOME/.kube/config
2020/11/24 00:34:35 Loading kubeconfig
2020/11/24 00:34:35 Deleted the resource <web> sucessfully, however dependent resources are not deleted
```

