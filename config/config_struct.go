package config

// Config
type Config struct {
	Github Github `yaml:"github"`
	Tidb   Tidb   `yaml:"tidb"`
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
	Redis  Redis  `yaml:"redis"`
	Pulsar Pulsar `yaml:"pulsar"`
}

// Tidb
type Tidb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Sql      Sql    `yaml:"sql"`
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

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Github
type Github struct {
	Loop   Loop     `yaml:"loop"`
	Tokens []string `yaml:"tokens"`
}

// Loop
type Loop struct {
	Break   int `yaml:"break"`
	Timeout int `yaml:"timeout"`
}

// Sql
type Sql struct {
	EventsDaily string `yaml:"eventsDaily"`
}

// Server
type Server struct {
	Port      int    `yaml:"port"`
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
}

// Redis
type Redis struct {
	Db       int    `yaml:"db"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

// Pulsar
type Pulsar struct {
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}
