# Example Service

This is the Example service

Generated with

```
micro new github.com/hb-chen/gateway/example/micro --namespace=go.micro --alias=example --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.example
- Type: srv
- Alias: example

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./example-srv
```

Build a docker image
```
make docker
```