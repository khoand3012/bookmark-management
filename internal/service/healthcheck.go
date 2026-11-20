package service

type healthCheck struct {
	serviceName string
	instanceId  string
}

type HealthCheckResult struct {
	Message     string `json:"message" example:"OK"`
	ServiceName string `json:"service_name" example:"bookmark_service"`
	InstanceId  string `json:"instance_id" example:"3f2b8c1e-7a54-4d0e-9b6a-1c2d3e4f5a6b"`
}

//go:generate mockery --name=HealthCheck --filename healthcheck.go --outpkg service_mocks
type HealthCheck interface {
	CheckHealth() *HealthCheckResult
}

func NewHealthCheck(serviceName string, instanceId string) HealthCheck {
	return &healthCheck{
		serviceName: serviceName,
		instanceId:  instanceId,
	}
}

func (h *healthCheck) CheckHealth() *HealthCheckResult {
	return &HealthCheckResult{
		Message:     "OK",
		ServiceName: h.serviceName,
		InstanceId:  h.instanceId,
	}
}
