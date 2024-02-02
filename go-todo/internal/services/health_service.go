package services

// ------------------------------STRUCTS---------------------------------- //

type HealthService struct{}

// ------------------------------METHODS---------------------------------- //

func (s *HealthService) Check() (*string, error) {
	msg := "Service health: OK"
	return &msg, nil
}
