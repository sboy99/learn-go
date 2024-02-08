package ports

// -------------------------------INTERFACES--------------------------------- //

type IJWTAdapterPort interface {
	Sign(paylod map[string]interface{}) (string, error)
}
