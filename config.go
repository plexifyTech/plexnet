package plexnet

type ApplicationConfig struct {
	Port     string `yaml:"port" env:"PORT"`
	AppName  string `yaml:"appName"`
	BasePath string `yaml:"basePath"`
}
