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
func GetInsecureClient(host string, port uint, config configs.GigamonoConfig) (interface{}, error) {
	conn, err := grpc.Dial(fmt.Sprint(host, ":", port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	switch port {
	case config.Services.API.Ports.Private:
		return generated.NewAPIClient(conn), nil
	case config.Services.Auth.Ports.Private:
		return generated.NewAuthClient(conn), nil
	case config.Services.WorkflowEngine.MainServer.Ports.Private:
		return generated.NewWorkflowMainServerClient(conn), nil
	default:
		return nil, errors.New("port is not recognised")
	}
}
