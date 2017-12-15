FROM golang

COPY target/go-hello-etcd /app/

WORKDIR /app
ENTRYPOINT ["./go-hello-etcd"]