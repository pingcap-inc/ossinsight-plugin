package config

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

// Interval
type Interval struct {
	YearCount string `yaml:"yearCount"`
	DayCount  string `yaml:"dayCount"`
	Daily     string `yaml:"daily"`
	Retry     int    `yaml:"retry"`
	RetryWait int    `yaml:"retryWait"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Tidb
type Tidb struct {
	Sql      Sql    `yaml:"sql"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

// Sql
type Sql struct {
	EventsDaily         string `yaml:"eventsDaily"`
	PrToday             string `yaml:"prToday"`
	PrDeveloperToday    string `yaml:"prDeveloperToday"`
	PrThisYear          string `yaml:"prThisYear"`
	PrDeveloperThisYear string `yaml:"prDeveloperThisYear"`
}

// Config
type Config struct {
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Redis    Redis    `yaml:"redis"`
	Pulsar   Pulsar   `yaml:"pulsar"`
	Github   Github   `yaml:"github"`
	Interval Interval `yaml:"interval"`
	Tidb     Tidb     `yaml:"tidb"`
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
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Consumer
type Consumer struct {
	Concurrency int    `yaml:"concurrency"`
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
}

