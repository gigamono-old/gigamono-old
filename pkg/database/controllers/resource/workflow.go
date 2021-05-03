package resource

import "github.com/gigamono/gigamono/pkg/database/models/resource"

// CreateWorkflow
func CreateWorkflow(workflow *resource.Workflow) (string, error) {
	return "", nil
}

func ActivateWorkflow(id *string) (string, error) {
	return "", nil
}

func GetWorkflow(id *string) (*resource.Workflow, error) {
	return &resource.Workflow{}, nil
}
