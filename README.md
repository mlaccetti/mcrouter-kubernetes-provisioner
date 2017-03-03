# mcrouter-k8s-provisioner

A mechanism to automatically configure Facebook's Mcrouter based on adding/removing memcached pods.

In simple terms, it'll automatically re-create the `mcrouter` config file for a bunch of `memcached` instances managed
through a Kubernetes deployment. The information for which `memcached` instances are available to put into the config is
extracted via DNS lookup in Kubernetes.

### Usage

Please see the `example` directory to see how this all works.

Pre-conditions:

- You have deployed the persistent volume
- You have deployed the persistent volume claim

First, deploy the `mcrouter-provisioner` - it'll mount the volume, create an empty config, and start listening for pod
creation/removal events. As `memcached` pods come online, it'll grab their info and update the `mcrouter` config file.
We are using a `deployment` as we want to ensure that one instance is always running.

Second, deploy a bunch of `memcached` nodes using a deployment and a headless service. We specifically want the headless
service so that we can look up all the nodes in Kubernetes' DNS.

Last, setup a Kubernetes deployment using `jamescarr/mcrouter` as the image, this time with a proper service. These
nodes all need a persistent volume to read the same configuration file from. When you build a service that requires
memcached, you will point it to this service, and one of the `mcrouter` instances will handle the request and route
to one of the `memcached` instances appropriately.

### Developing

Make sure you have the `dep` tool installed:

```
go get -u github.com/golang/dep/...
```

Next, make sure you install the dependencies we need:

```
dep ensure
```

Finally, build the app:

```
go build -o mcrouter-k8s-provisioner main.go
```
