package config

// Server
type Server struct {
	Port   int    `yaml:"port"`
	Health string `yaml:"health"`
}

// Redis
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// Pulsar
type Pulsar struct {
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Config
type Config struct {
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
	Redis  Redis  `yaml:"redis"`
	Pulsar Pulsar `yaml:"pulsar"`
	Github Github `yaml:"github"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Producer
type Producer struct {
	Topic string `yaml:"topic"`
	Retry int    `yaml:"retry"`
}

// Github
type Github struct {
	Loop   Loop     `yaml:"loop"`
	Tokens []string `yaml:"tokens"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

