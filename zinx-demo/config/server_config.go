package config

type ServerConfig struct {
	Server Server `yaml:"server"`
	Logrus Logrus `yaml:"logrus"`
}

type Server struct {
	Ip               string `yaml:"ip"`
	Port             int    `yaml:"port"`
	HandleWorkerSize uint32 `yaml:"handleWorkerSize"`
	MaxWaitSize      uint32 `yaml:"maxWaitSize"`
}

type Logrus struct {
	MaxCount int    `yaml:"maxCount"`
	FileName string `yaml:"fileName"`
	MaxSize  int    `yaml:"maxSize"`
}
