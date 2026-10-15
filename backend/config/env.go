package config

type Config struct {
	AppPort string
	DbHost  string
}

func Load() (*Config, error) {

}
