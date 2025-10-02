package config

import (
	"time"
)

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

	RabbitMQ struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"rabbitmq"`

	Kafka struct {
		Brokers           string        `mapstructure:"brokers"`
		BlocksTopicPrefix string        `mapstructure:"blocks_topic_prefix"`
		MaxAttempts       int           `mapstructure:"max_attempts"`
		MessageMaxBytes   int           `mapstructure:"message_max_bytes"`
		RetentionTime     time.Duration `mapstructure:"retention_time"`
		Partitions        int           `mapstructure:"partitions"`
		ReplicationFactor int           `mapstructure:"replication_factor"`
	} `mapstructure:"kafka"`

	Jwt struct {
		Secret          string `mapstructure:"secret"`
		TimeDurationJwt string `mapstructure:"time_duration_jwt"`
	} `mapstructure:"jwt"`

	Email struct {
		Username string `mapstructure:"user_name"`
		Password string `mapstructure:"password"`
		Port     int    `mapstructure:"port"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"email"`

	Redis struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"redis"`

	Google struct {
		ClientId string `mapstructure:"client_id"`
	} `mapstructure:"google"`

	Apple struct {
		AppIdIos     string `mapstructure:"app_id_ios"`
		AppIdAndroid string `mapstructure:"app_id_android"`
		TeamId       string `mapstructure:"team_id"`
		KeyId        string `mapstructure:"key_id"`
		PathAuthKey  string `mapstructure:"path_auth_key"`
	} `mapstructure:"apple"`

	Api struct {
		EOACipherKey string `mapstructure:"eoa_cipher_key"`
		PaymentGate  struct {
			ApiKey        string `mapstructure:"api_key"`
			ApiSecret     string `mapstructure:"api_secret"`
			ApiBaseUrl    string `mapstructure:"api_base_url"`
			EncryptionKey string `mapstructure:"encryption_key"`
		} `mapstructure:"payment_gate"`
		FactoryOwnerPrivateKey       string `mapstructure:"factory_owner_private_key"`
		VetifyingPaymasterPrivateKey string `mapstructure:"verifying_paymaster_private_key"`
		ReservationTime              string `mapstructure:"reservation_time"`
	} `mapstructure:"api"`

	Scheduler struct {
		Interval        time.Duration `mapstructure:"interval"`
		BackoffInterval time.Duration `mapstructure:"backoff_interval"`
		CoingeckoUrl    string        `mapstructure:"coingecko_url"`
	} `mapstructure:"scheduler"`

	SmartAccount struct {
		Bin string `mapstructure:"bin"`
	} `mapstructure:"smartAccount"`

	RelayerPrivateKey struct {
		AccountPrivateKey string `mapstructure:"account_private_key"`
	} `mapstructure:"relayerPrivateKey"`

	TimeDurationCode string `mapstructure:"timeDurationVerifyCode"`

	Shopify struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"shopify"`

	Chain struct {
		Rpc                                   string `mapstructure:"rpc"`
		PrivateKey                            string `mapstructure:"private_key"`
		BundlerAvaxERC4337BundlerETHClientURL string `mapstructure:"bundler_avax_erc4337_bundler_eth_client_url"`
		HexKeySystem                          string `mapstructure:"hex_key_system"`
	} `mapstructure:"chain"`

	Indexer struct {
		UrlDefault string `mapstructure:"url_default"`
		UrlSandBox string `mapstructure:"url_sandbox"`
	} `mapstructure:"indexer"`

	WalletService struct {
		Dest   string `mapstructure:"dest"`
		Port   string `mapstructure:"port"`
		Secret string `mapstructure:"secret"`
	} `mapstructure:"walletService"`

	Moralis struct {
		APIKey  string `mapstructure:"api_key" yaml:"api_key"`
		BaseURL string `mapstructure:"base_url" yaml:"base_url"`
		Chain   string `mapstructure:"chain" yaml:"chain"`
	} `mapstructure:"moralis" yaml:"moralis"`
}
