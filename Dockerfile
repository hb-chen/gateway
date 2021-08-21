FROM alpine:3.10

ADD bin/linux/gateway /opt/grpc-gateway/gateway

EXPOSE 8080
WORKDIR /opt/grpc-gateway

ENTRYPOINT [ "./gateway", "-e" ]
