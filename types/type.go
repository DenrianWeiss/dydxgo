package types

type ApiKeyCredentials struct {
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

type RateLimit struct {
	Remaining  string
	Reset      string
	RetryAfter string
	Limit      string
}
