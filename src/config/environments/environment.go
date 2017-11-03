package environments

import "os"

const envVarName = "CHIP_ENV"

func appEnvironment() string {
	return os.Getenv(envVarName)
}
