package inits

// ServiceKind represents the different types of Gigamono services.
type ServiceKind string

// ...
const (
	API                              ServiceKind = "API"
	Auth                             ServiceKind = "Auth"
	WorkflowEngineMainServer         ServiceKind = "WorkflowEngineMainServer"
	WorkflowEngineAPIService         ServiceKind = "WorkflowEngineAPIService"
	WorkflowEngineRunnableSupervisor ServiceKind = "WorkflowEngineRunnableSupervisor"
	BaseEngineMainServer             ServiceKind = "BaseEngineMainServer"
)

// DatabaseKind gets the database kind supported by service.
func (kind *ServiceKind) DatabaseKind() string {
	switch *kind {
	case API,
		WorkflowEngineMainServer,
		WorkflowEngineAPIService,
		WorkflowEngineRunnableSupervisor,
		BaseEngineMainServer:
		return "resource"
	case Auth:
		return "auth"
	}

	return ""
}
