package entity

// DefaultCredentials represents email/password combination.
type DefaultCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
