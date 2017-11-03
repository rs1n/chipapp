package environments

import "os"

const envVarName = "CHIPAPP_ENV"

func appEnvironment() string {
	return os.Getenv(envVarName)
}
