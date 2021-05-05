package grpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/services/proto/generated"
)

// GetInsecureClient returns a connection interface of supported gRPC client.
// Sec: The assumption is that the services will run together in TLS-protected Kubernetes cluster.
// TODO: Sec: Support optional usage of TLS cert. config.Services.TLS.
func GetInsecureClient(host string, port int, config configs.GigamonoConfig) (interface{}, error) {
	conn, err := grpc.Dial(fmt.Sprint(host, ":", port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	switch port {
	case config.Services.Types.Auth.Port:
		return generated.NewAuthClient(conn), nil
	case config.Services.Types.WorkflowEngine.Ports.MainServer:
		return generated.NewWorkflowMainServerClient(conn), nil
	case config.Services.Types.API.Port:
		return generated.NewAPIClient(conn), nil
	default:
		return nil, errors.New("port is not recognised")
	}
}
