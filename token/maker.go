package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for specific username & duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken checks if a token is valid or not
	VerifyToken(token string) (*Payload, error)
}
