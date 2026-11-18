package service

type healthCheck struct {
	serviceName string
	instanceId  string
}

type HealthCheckResult struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceId  string `json:"instance_id"`
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
