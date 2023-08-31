package types

type ApiKeyCredentials struct {
	Key        string
	Secret     string
	Passphrase string
}

type RateLimit struct {
	Remaining  string
	Reset      string
	RetryAfter string
	Limit      string
}
