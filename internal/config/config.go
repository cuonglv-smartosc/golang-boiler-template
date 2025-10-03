package config

// Default is a config instance.
var Default Config //nolint:gochecknoglobals // config must be global

type Config struct {
	LogLevel string `mapstructure:"log_level"`

	Port string `mapstructure:"port"`

	Gin struct {
		Mode string `mapstructure:"mode"`
	} `mapstructure:"gin"`

	Swagger struct {
		Hostname string `mapstructure:"hostname"`
	} `mapstructure:"swagger"`

	Sentry struct {
		DSN        string  `mapstructure:"dsn"`
		SampleRate float32 `mapstructure:"sample_rate"`
	} `mapstructure:"sentry"`

	Database struct {
		URL string `mapstructure:"url"`
		Log bool   `mapstructure:"log"`
	} `mapstructure:"database"`

	Jwt struct {
		Secret          string `mapstructure:"secret"`
		TimeDurationJwt string `mapstructure:"time_duration_jwt"`
	} `mapstructure:"auth"`

	Redis struct {
		URL      string `mapstructure:"url"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`
}
