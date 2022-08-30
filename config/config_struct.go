package config

// Disable
type Disable struct {
	Producer bool `yaml:"producer"`
	Interval bool `yaml:"interval"`
}

// Producer
type Producer struct {
	Retry int    `yaml:"retry"`
	Topic string `yaml:"topic"`
}

// Consumer
type Consumer struct {
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
	Topic       string `yaml:"topic"`
}

// Api
type Api struct {
	Version int `yaml:"version"`
}

// Tidb
type Tidb struct {
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Sql      Sql    `yaml:"sql"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
}

// Server
type Server struct {
	Port      int    `yaml:"port"`
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
}

// Pulsar
type Pulsar struct {
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
}

// Config
type Config struct {
	Disable  Disable  `yaml:"disable"`
	Server   Server   `yaml:"server"`
	Pulsar   Pulsar   `yaml:"pulsar"`
	Github   Github   `yaml:"github"`
	Interval Interval `yaml:"interval"`
	Api      Api      `yaml:"api"`
	Redis    Redis    `yaml:"redis"`
	Tidb     Tidb     `yaml:"tidb"`
	Log      Log      `yaml:"log"`
}

// Interval
type Interval struct {
	Daily     string `yaml:"daily"`
	Retry     int    `yaml:"retry"`
	RetryWait int    `yaml:"retryWait"`
}

// Redis
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// Sql
type Sql struct {
	EventsDaily         string `yaml:"eventsDaily"`
	PrDaily             string `yaml:"prDaily"`
	PrDeveloperDaily    string `yaml:"prDeveloperDaily"`
	PrDeveloperThisYear string `yaml:"prDeveloperThisYear"`
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

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

