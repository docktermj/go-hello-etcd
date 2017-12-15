# go-hello-etcd

## Usage

A simple program to show how to integrate
[etcd](https://coreos.com/etcd/).

### Invocation

```console
go-hello-etcd
```

## Demonstrate

Verify `docker` network is 172.17.0.1.
If gateway is not 172.17.0.1, the following `docker` statements need to be modified before being run.

```console
docker network inspect bridge | grep Gateway
```

Build etcd and docker demo images.

In terminal #1

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-hello-etcd"

cd ${REPOSITORY_DIR}
make dependencies
```

Fix [bug](https://github.com/coreos/etcd/issues/8715).

```console
rm  ${REPOSITORY_DIR}/vendor/github.com/coreos/etcd/client/keys.generated.go
```

Make docker image

```console
cd ${REPOSITORY_DIR}
make build-demo
```

Clean old docker images

```console
docker rm etcd2 etcd3 etcd4
```

Start "local" etcd

```console
cd ~
rm -rf ~/default.etcd

export ETCD_HOST=172.17.0.1
go-hello-etcd \
  --name etcd1 \
  --advertise-client-urls http://${ETCD_HOST}:2379,http://${ETCD_HOST}:4001 \
  --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
  --initial-advertise-peer-urls http://${ETCD_HOST}:2380 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster-state new \
  --initial-cluster "etcd1=http://172.17.0.1:2380,etcd2=http://172.17.0.2:23802,etcd3=http://172.17.0.3:23803,etcd4=http://172.17.0.4:23804"
```

In terminal #2

```console
export ETCD_HOST=172.17.0.2
docker run \
  --name etcd2 \
  --env ETCD_NAME=etcd2 \
  --env ETCD_ADVERTISE_CLIENT_URLS=http://${ETCD_HOST}:23792,http://${ETCD_HOST}:40012 \
  --env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:23792,http://0.0.0.0:40012 \
  --env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${ETCD_HOST}:23802 \
  --env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:23802 \
  --env ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster-1 \
  --env ETCD_INITIAL_CLUSTER_STATE=new \
  --env ETCD_INITIAL_CLUSTER="etcd1=http://172.17.0.1:2380,etcd2=http://172.17.0.2:23802,etcd3=http://172.17.0.3:23803,etcd4=http://172.17.0.4:23804" \
  --publish 23792:2379 \
  --publish 23802:2380 \
  --publish 40012:4001 \
  local/go-hello-etcd-demo
```

In terminal #3, having IP address 172.17.0.3

```console
export ETCD_HOST=172.17.0.3
docker run \
  --name etcd3 \
  --env ETCD_NAME=etcd3 \
  --env ETCD_ADVERTISE_CLIENT_URLS=http://${ETCD_HOST}:23793,http://${ETCD_HOST}:40013 \
  --env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:23793,http://0.0.0.0:40013 \
  --env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${ETCD_HOST}:23803 \
  --env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:23803 \
  --env ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster-1 \
  --env ETCD_INITIAL_CLUSTER_STATE=new \
  --env ETCD_INITIAL_CLUSTER="etcd1=http://172.17.0.1:2380,etcd2=http://172.17.0.2:23802,etcd3=http://172.17.0.3:23803,etcd4=http://172.17.0.4:23804" \
  --publish 23793:2379 \
  --publish 23803:2380 \
  --publish 40013:4001 \
  local/go-hello-etcd-demo
```

In terminal #4, having IP address 172.17.0.4

```console
export ETCD_HOST=172.17.0.4
docker run \
  --name etcd4 \
  --env ETCD_NAME=etcd4 \
  --env ETCD_ADVERTISE_CLIENT_URLS=http://${ETCD_HOST}:23794,http://${ETCD_HOST}:40014 \
  --env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:23794,http://0.0.0.0:40014 \
  --env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${ETCD_HOST}:23804 \
  --env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:23804 \
  --env ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster-1 \
  --env ETCD_INITIAL_CLUSTER_STATE=new \
  --env ETCD_INITIAL_CLUSTER="etcd1=http://172.17.0.1:2380,etcd2=http://172.17.0.2:23802,etcd3=http://172.17.0.3:23803,etcd4=http://172.17.0.4:23804" \
  --publish 23794:2379 \
  --publish 23804:2380 \
  --publish 40014:4001 \
  local/go-hello-etcd-demo
```

In terminal #5, try these commands

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"

docker cp ${GOPATH}/bin/etcdctl etcd2:/app
docker exec -it etcd2 bash

export ETCDCTL_API=3
etcdctl member list
etcdctl get bob
etcdctl put bob 7
etcdctl get bob
```

In terminal #6, try these commands

```console
docker exec -it etcd2 bash
```

In docker container, try these commands

```console
export ETCDCTL_API=3
/app/etcdctl --endpoints "http://127.0.0.1:23792,http://127.0.0.1:40012" member list
/app/etcdctl --endpoints "http://127.0.0.1:23792,http://127.0.0.1:40012" get bob
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-hello-etcd"
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-hello-etcd.git
```

#### Download dependencies

```console
cd ${REPOSITORY_DIR}
make dependencies
```

Etcd command-line tool

```console
go get -u github.com/coreos/etcd/etcdctl
```

### Build

#### Local build

```console
cd ${REPOSITORY_DIR}
make
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${REPOSITORY_DIR}
make build
```

The results will be in the `.../target` directory.

### Test

```console
cd ${REPOSITORY_DIR}
make test-local
```

### Cleanup

```console
cd ${REPOSITORY_DIR}
make clean
```

### Referenes
