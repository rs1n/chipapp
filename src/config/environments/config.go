package environments

const (
	envProduction = "production"
	envTest       = "test"
)

type Config struct {
	IsDebug bool
	Port    int
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
