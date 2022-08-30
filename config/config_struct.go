package config

// Config
type Config struct {
	Api      Api      `yaml:"api"`
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Pulsar   Pulsar   `yaml:"pulsar"`
	Interval Interval `yaml:"interval"`
	Disable  Disable  `yaml:"disable"`
	Redis    Redis    `yaml:"redis"`
	Github   Github   `yaml:"github"`
	Tidb     Tidb     `yaml:"tidb"`
}

// Interval
type Interval struct {
	Retry     int    `yaml:"retry"`
	RetryWait int    `yaml:"retryWait"`
	YearCount string `yaml:"yearCount"`
	DayCount  string `yaml:"dayCount"`
	Daily     string `yaml:"daily"`
}

// Sql
type Sql struct {
	EventsDaily         string `yaml:"eventsDaily"`
	PrToday             string `yaml:"prToday"`
	PrDeveloperToday    string `yaml:"prDeveloperToday"`
	PrThisYear          string `yaml:"prThisYear"`
	PrDeveloperThisYear string `yaml:"prDeveloperThisYear"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Disable
type Disable struct {
	Interval bool `yaml:"interval"`
	Producer bool `yaml:"producer"`
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

// Api
type Api struct {
	Version int `yaml:"version"`
}

// Server
type Server struct {
	Port      int    `yaml:"port"`
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Pulsar
type Pulsar struct {
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
	Host     string   `yaml:"host"`
}

// Producer
type Producer struct {
	Topic string `yaml:"topic"`
	Retry int    `yaml:"retry"`
}

// Redis
type Redis struct {
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Host     string `yaml:"host"`
}

// Tidb
type Tidb struct {
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Sql      Sql    `yaml:"sql"`
	Host     string `yaml:"host"`
}

