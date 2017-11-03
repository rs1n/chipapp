package environments

const (
	envProduction = "production"
	envTest       = "test"
)

type Config struct {
	Port int
}

func GetConfig() Config {
	switch appEnvironment() {
	case envProduction:
		return production
	case envTest:
		return test
	}
	return development // Default configuration.
}
