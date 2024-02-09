package ports

// -------------------------------STRUCTS--------------------------------- //

type AccessTokenPayload struct {
	Subject   string `json:"sub"`
	UserId    uint   `json:"userId"`
	Name      string `json:"name"`
	ExpiresAt int64  `json:"exp"`
}

// -------------------------------INTERFACES--------------------------------- //

type IJwtExtractorStrategyPort interface {
	Extract() (string, error)
}

type IJWTAdapterPort interface {
	Sign(paylod interface{}) (string, error)
	Validate(token string, payload interface{}) error
}
