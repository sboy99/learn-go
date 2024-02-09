package stratigies

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ------------------------------STRUCTS---------------------------------- //

type BearerExtractorStrategy struct {
	Ctx *fiber.Ctx
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (s *BearerExtractorStrategy) Extract() (string, error) {
	bearerToken := s.Ctx.Get("Authorization")

	if s.isEmptyBearerToken(bearerToken) {
		// If Authorization header is missing, return an error response
		return "", errors.New("Authorization header is missing")
	}

	// Check if the Authorization header starts with "Bearer"
	if s.isInvalidBearerTokenFormat(bearerToken) {
		// If Authorization header does not start with "Bearer", return an error response
		return "", errors.New("Invalid Authorization header format")
	}

	// Extract the token from the Authorization header
	token := s.extractToken(bearerToken)

	// Done
	return token, nil
}

// ------------------------------PRIVATE_METHODS---------------------------------- //

func (s *BearerExtractorStrategy) isEmptyBearerToken(bearerToken string) bool {
	return bearerToken == ""
}

func (s *BearerExtractorStrategy) isInvalidBearerTokenFormat(bearerToken string) bool {
	return !strings.HasPrefix(bearerToken, "Bearer ")
}

func (s *BearerExtractorStrategy) extractToken(bearerToken string) string {
	return strings.TrimPrefix(bearerToken, "Bearer ")
}
