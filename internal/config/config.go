package config

type Config struct {
	Workers   int
	RateLimit int
	UserAgent string
}

func Default() Config {
	return Config{
		Workers:   5,
		RateLimit: 10,
		UserAgent: "Gofuzzer/0.1",
	}
}
