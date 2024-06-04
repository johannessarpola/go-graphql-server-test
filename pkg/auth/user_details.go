package auth

type UserDetails struct {
	Login         string `json:"login"`
	Authenticated bool   `json:"authenticated"`
}
