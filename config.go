package plexnet

var (
	config ApplicationConfig
)

type ApplicationConfig struct {
	Port     string `yaml:"port" env:"PORT"`
	AppName  string `yaml:"appName"`
	BasePath string `yaml:"basePath"`
}

func Init(c ApplicationConfig) {
	config = c
}
