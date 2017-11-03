package environment

var environment *Environment

func GetEnvironment() *Environment {
	if environment == nil {
		panic("environment.GetEnvironment: environment is not initialized")
	}
	return environment
}

func InitializeEnvironmentFor(rhp HtmlRendererParams) {
	if environment != nil {
		panic("environment.InitializeEnvironmentFor: environment is already initialized")
	}
	environment = NewEnvironment(rhp)
}
