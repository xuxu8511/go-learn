package config

type ServerConfig struct {
	Server Server `yaml:"server"`
	Logrus Logrus `yaml:"logrus"`
}

type Server struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Logrus struct {
	Filename string `yaml:"filename"`
	Maxsize  int    `yaml:"maxsize"`
	Maxcount int    `yaml:"maxcount"`
}
