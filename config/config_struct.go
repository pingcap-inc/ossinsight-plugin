package config

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Config
type Config struct {
	Http   Http   `yaml:"http"`
	Log    Log    `yaml:"log"`
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

// Pulsar
type Pulsar struct {
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Producer
type Producer struct {
	Topic string `yaml:"topic"`
	Retry int    `yaml:"retry"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Github
type Github struct {
	Loop Loop `yaml:"loop"`
}

