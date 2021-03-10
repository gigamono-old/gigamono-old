package services

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/sageflow/sageflow/pkg/services/proto"
	"github.com/sageflow/sageflow/pkg/configs"
)

// GetInsecureServiceClient returns a connection interface of supported gRPC client.
// Sec: The assumption is that the servers will run in the same cluster so HTTPS connnection is not important.
func GetInsecureServiceClient(host string, port int, config configs.SageflowConfig) (interface{}, error) {
	conn, err := grpc.Dial(fmt.Sprint(host, ":", port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	switch port {
	case config.Server.Auth.Port:
		return proto.NewAuthServiceClient(conn), nil
	case config.Server.Engine.Port:
		return proto.NewEngineServiceClient(conn), nil
	case config.Server.API.Port:
		return proto.NewAPIServiceClient(conn), nil
	default:
		return nil, errors.New("Port is not recognised")
	}
}
