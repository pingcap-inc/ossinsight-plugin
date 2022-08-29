package config

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

// Sql
type Sql struct {
	EventsDaily         string `yaml:"eventsDaily"`
	PrToday             string `yaml:"prToday"`
	PrDeveloperToday    string `yaml:"prDeveloperToday"`
	PrThisYear          string `yaml:"prThisYear"`
	PrDeveloperThisYear string `yaml:"prDeveloperThisYear"`
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

// Redis
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// Producer
type Producer struct {
	Retry int    `yaml:"retry"`
	Topic string `yaml:"topic"`
}

// Consumer
type Consumer struct {
	Concurrency int    `yaml:"concurrency"`
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
}

// Interval
type Interval struct {
	RetryWait int    `yaml:"retryWait"`
	YearCount string `yaml:"yearCount"`
	DayCount  string `yaml:"dayCount"`
	Daily     string `yaml:"daily"`
	Retry     int    `yaml:"retry"`
}

// Tidb
type Tidb struct {
	Db       string `yaml:"db"`
	Sql      Sql    `yaml:"sql"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Config
type Config struct {
	Pulsar   Pulsar   `yaml:"pulsar"`
	Github   Github   `yaml:"github"`
	Interval Interval `yaml:"interval"`
	Tidb     Tidb     `yaml:"tidb"`
	Api      Api      `yaml:"api"`
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Redis    Redis    `yaml:"redis"`
}

// Pulsar
type Pulsar struct {
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
	Host     string   `yaml:"host"`
}

