## Demo project to demonstrate persistent connections in k8s

Run to start
```shell
make all
```

Observe traffic distribution with
```shell
kubectl logs -f deployment/client
```

To play with istio run
```shell
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.24.2 sh -
./istio-1.24.2/bin/istioctl install -y
kubectl label namespace default istio-injection=enabled
```
When istio is ready (all pods in istio-system namespace are running)
```shell
kubectl scale deployment server --replicas 0
kubectl scale deployment server --replicas 3
kubectl scale deployment client --replicas 0
kubectl scale deployment client --replicas 1
```

Verify envoy-proxy is attached to pods (2/2 READY)
```shell
kubectl get po
NAME                      READY   STATUS    RESTARTS   AGE
client-5db8dfb89f-wgv54   2/2     Running   0          3s
server-7f78cdb968-6pnsm   2/2     Running   0          4s
server-7f78cdb968-hpr7n   2/2     Running   0          4s
server-7f78cdb968-vhj47   2/2     Running   0          4s
```

Then make changes to k8s/istio/vs-server.yaml and apply to see how it affects traffic
```shell
kubectl apply -f k8s/istio
```

Try to use DestinationRule
1. Without trafficPolicy section. You will observe "random" distribution
2. With trafficPolicy section and `useSourceIp: true`. You will observe "sticky" requests.
3. With trafficPolicy section and `httpHeaderName: x-user`. You will observe "sticky" requests based on header value.