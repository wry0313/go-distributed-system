package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
  RequiredServices []ServiceName
  ServiceUpdateUrl string
  HeartbeatURL string
}

type ServiceName string

const (
	LogService = ServiceName("Log Service")
	GradingService = ServiceName("Grading Service")
  PortalService = ServiceName("Portal Service")
)

type patchEntry struct {
  Name ServiceName
  URL string
}

type patch struct {
  Added []patchEntry
  Removed []patchEntry
}
