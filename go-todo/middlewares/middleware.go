package middlewares

import (
	"github.com/sboy99/learn-go/go-todo/internal/ports"
)

// ------------------------------STRUCTS---------------------------------- //

type Middleware struct {
	JwtAdapter ports.IJWTAdapterPort
}
