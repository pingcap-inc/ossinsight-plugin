package config

// Pulsar
type Pulsar struct {
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Redis
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
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

// Config
type Config struct {
	Http   Http   `yaml:"http"`
	Log    Log    `yaml:"log"`
	Redis  Redis  `yaml:"redis"`
	Pulsar Pulsar `yaml:"pulsar"`
	Github Github `yaml:"github"`
}

// Http
type Http struct {
	Port int `yaml:"port"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

