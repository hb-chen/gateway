package util

import (
	"github.com/micro/go-micro/v2/util/addr"
	mNet "github.com/micro/go-micro/v2/util/net"
	"net"
	"strings"
)

func Address(address string) (string, error) {
	var err error
	var host, port string

	if cnt := strings.Count(address, ":"); cnt >= 1 {
		// ipv6 address in format [host]:port or ipv4 host:port
		host, port, err = net.SplitHostPort(address)
		if err != nil {
			return "", err
		}
	} else {
		host = address
	}

	addr, err := addr.Extract(host)
	if err != nil {
		return "", err
	}

	// mq-rpc(eg. nats) doesn't need the port. its addr is queue name.
	if port != "" {
		addr = mNet.HostPort(addr, port)
	}

	return addr, nil
}
