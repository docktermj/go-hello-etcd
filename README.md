# go-hello-etcd

## Usage

A simple program to show how to integrate
[etcd](https://coreos.com/etcd/).

### Invocation

```console
go-hello-etcd
```

## Demonstrate

In terminal #1

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-hello-etcd"

cd ${REPOSITORY_DIR}
make dependencies

cd ${REPOSITORY_DIR}
make build-demo
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


