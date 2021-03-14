package services

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/sageflow/sageflow/pkg/services/proto/generated"
	"github.com/sageflow/sageflow/pkg/configs"
)

// GetInsecureServiceClient returns a connection interface of supported gRPC client.
// Sec: The assumption is that the services will run together in TLS-protected Kubernetes cluster.
// TODO: Support optional TLS cert. config.
func GetInsecureServiceClient(host string, port int, config configs.SageflowConfig) (interface{}, error) {
	conn, err := grpc.Dial(fmt.Sprint(host, ":", port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	switch port {
	case config.Services.Types.Auth.Port:
		return generated.NewAuthServiceClient(conn), nil
	case config.Services.Types.Engine.Port:
		return generated.NewEngineServiceClient(conn), nil
	case config.Services.Types.API.Port:
		return generated.NewAPIServiceClient(conn), nil
	default:
		return nil, errors.New("port is not recognised")
	}
}
