package spotify

type AuthConfig struct {
	ClientID      string   `yaml:"client_id"`
	ClientSecret  string   `yaml:"client_secret"`
	TokenEndpoint string   `yaml:"token_endpoint"`
	AuthEndpoint  string   `yaml:"auth_endpoint"`
	RedirectURL   string   `yaml:"redirect_url"`
	Scopes        []string `yaml:"scopes"`
}

type Config struct {
	Auth AuthConfig `yaml:"auth"`
	Base string     `yaml:"base"`
}
