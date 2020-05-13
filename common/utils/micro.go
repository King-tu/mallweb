package utils

import (
	"time"

	"github.com/micro/go-micro/v2/client"
)

const (
	defaultTimeout = time.Second * 360
)

func RPCClient() client.Client {
	return client.NewClient(client.RequestTimeout(defaultTimeout))
}

func ClientWithTimeout(seconds int) client.Client {
	timeout := time.Second * time.Duration(seconds)
	return client.NewClient(client.RequestTimeout(timeout))
}

func VideoTransRPCClient() client.Client {
	return client.NewClient(client.RequestTimeout(time.Second * 3600))
}
