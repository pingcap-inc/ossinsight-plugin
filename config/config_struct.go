package config

// Config
type Config struct {
	Pulsar Pulsar `yaml:"pulsar"`
	Http   Http   `yaml:"http"`
	Log    Log    `yaml:"log"`
}

// Pulsar
type Pulsar struct {
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
}

// Producer
type Producer struct {
	Topic string `yaml:"topic"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
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

