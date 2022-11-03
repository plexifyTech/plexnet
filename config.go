package plexnet

type AppConfig struct {
	Port     string `yaml:"port" env:"PORT"`
	AppName  string `yaml:"appName" env:"APP_NAME"`
	BasePath string `yaml:"basePath"`
}
