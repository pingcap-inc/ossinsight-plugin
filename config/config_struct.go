package config

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Config
type Config struct {
	Disable  Disable  `yaml:"disable"`
	Log      Log      `yaml:"log"`
	Github   Github   `yaml:"github"`
	Api      Api      `yaml:"api"`
	Server   Server   `yaml:"server"`
	Redis    Redis    `yaml:"redis"`
	Pulsar   Pulsar   `yaml:"pulsar"`
	Interval Interval `yaml:"interval"`
	Tidb     Tidb     `yaml:"tidb"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Server
type Server struct {
	Port      int    `yaml:"port"`
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
}

// Pulsar
type Pulsar struct {
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Api
type Api struct {
	Version int `yaml:"version"`
}

// Interval
type Interval struct {
	Daily     string `yaml:"daily"`
	Language  string `yaml:"language"`
	Retry     int    `yaml:"retry"`
	RetryWait int    `yaml:"retryWait"`
}

// Disable
type Disable struct {
	Producer bool `yaml:"producer"`
	Interval bool `yaml:"interval"`
}

// Github
type Github struct {
	Loop   Loop     `yaml:"loop"`
	Tokens []string `yaml:"tokens"`
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

// Tidb
type Tidb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Sql      Sql    `yaml:"sql"`
}

// Sql
type Sql struct {
	PrDaily             string `yaml:"prDaily"`
	PrDeveloperDaily    string `yaml:"prDeveloperDaily"`
	PrDeveloperThisYear string `yaml:"prDeveloperThisYear"`
	LanguageToday       string `yaml:"languageToday"`
	EventsDaily         string `yaml:"eventsDaily"`
}

