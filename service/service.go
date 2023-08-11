package service

import (
	"context"
	"net/http"
)

func Start(
	ctx context.Context,
	serviceName, host, port string,
	registerHandlerFunc func(),
) (context.Context, error) {
	registerHandlerFunc()
	ctx = startService(ctx, serviceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, srviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
}
