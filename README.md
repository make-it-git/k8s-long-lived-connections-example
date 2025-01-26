## Demo project to demonstrate persistent connections in k8s

Run to start
```shell
make all
```

To play with istio run
```shell
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.24.2 sh -
./istio-1.24.2/bin/istioctl install -y
kubectl label namespace default istio-injection=enabled
```
And when istio is ready
```shell
kubectl scale deployment server --replicas 0
kubectl scale deployment server --replicas 3
kubectl scale deployment client --replicas 0
kubectl scale deployment client --replicas 1
```

Then make changes to k8s/istio/vs-server.yaml and apply to see how it affects traffic
```shell
kubectl apply -f k8s/istio
```