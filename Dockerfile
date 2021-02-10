FROM alpine:3.10

ADD bin/gateway /bin/gateway

WORKDIR /bin
ENTRYPOINT [ "gateway" ]
