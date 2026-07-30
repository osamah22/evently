package contracts

type MeResponse struct {
	Subject   string   `json:"subject"`
	Issuer    string   `json:"issuer"`
	Audience  []string `json:"audience"`
	Scope     string   `json:"scope,omitempty"`
	IssuedAt  int64    `json:"issued_at"`
	ExpiresAt int64    `json:"expires_at"`
}
