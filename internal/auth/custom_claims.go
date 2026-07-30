package auth

import "context"

type customClaims struct {
	Permissions []string `json:"permissions"`
}

func (c *customClaims) Validate(ctx context.Context) error {
	return nil
}
