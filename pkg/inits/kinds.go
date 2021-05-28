package inits

// ServiceKind represents the different types of Gigamono services.
type ServiceKind string

// ...
const (
	API                              ServiceKind = "API"
	Auth                             ServiceKind = "Auth"
	WorkflowEngineMainServer         ServiceKind = "WorkflowEngineMainServer"
	WorkflowEngineWebhookService     ServiceKind = "WorkflowEngineWebhookService"
	WorkflowEngineRunnableSupervisor ServiceKind = "WorkflowEngineRunnableSupervisor"
	DocumentEngineMainServer         ServiceKind = "DocumentEngineMainServer"
)

// DatabaseKind gets the database kind supported by service.
func (kind *ServiceKind) DatabaseKind() string {
	switch *kind {
	case API,
		WorkflowEngineMainServer,
		WorkflowEngineWebhookService,
		WorkflowEngineRunnableSupervisor,
		DocumentEngineMainServer:
		return "resource"
	case Auth:
		return "auth"
	}

	return ""
}
