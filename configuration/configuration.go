package configuration

type Configuration struct {
	RestConfiguration RestConfiguration
	OpenFeature       OpenFeature
}

type RestConfiguration struct {
	Host string `env:"REST_HOST, default=localhost"`
	Port string `env:"REST_PORT, default=8080"`
}
type OpenFeature struct {
	Client string `env:"OPEN_FEATURE_CLIENT, default=web-scrapper-api"`
	Host   string `env:"OPEN_FEATURE_HOST, default=localhost"`
	Port   int16  `env:"OPEN_FEATURE_PORT, default=8013"`
}
