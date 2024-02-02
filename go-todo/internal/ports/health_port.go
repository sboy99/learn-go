package ports

// ------------------------------INTERFACES---------------------------------- //

type IHealthServicePort interface {
	Check() (*string, error)
}
