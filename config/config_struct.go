package config

// Server
type Server struct {
	Port      int    `yaml:"port"`
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
}

// Redis
type Redis struct {
	Db       int    `yaml:"db"`
	Lua      Lua    `yaml:"lua"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Interval
type Interval struct {
	RetryWait    int    `yaml:"retryWait"`
	Daily        string `yaml:"daily"`
	Language     string `yaml:"language"`
	Latest       string `yaml:"latest"`
	LatestDuring int    `yaml:"latestDuring"`
	Retry        int    `yaml:"retry"`
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

// Pulsar
type Pulsar struct {
	Env      string   `yaml:"env"`
	DevHost  string   `yaml:"devHost"`
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Sql
type Sql struct {
	LanguageToday       string `yaml:"languageToday"`
	EventsDaily         string `yaml:"eventsDaily"`
	PrDaily             string `yaml:"prDaily"`
	PrDeveloperDaily    string `yaml:"prDeveloperDaily"`
	PrDeveloperThisYear string `yaml:"prDeveloperThisYear"`
}

// Lua
type Lua struct {
	MergeLatest string `yaml:"mergeLatest"`
}

// Disable
type Disable struct {
	Interval bool `yaml:"interval"`
	Producer bool `yaml:"producer"`
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

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Config
type Config struct {
	Disable  Disable  `yaml:"disable"`
	Pulsar   Pulsar   `yaml:"pulsar"`
	Github   Github   `yaml:"github"`
	Interval Interval `yaml:"interval"`
	Tidb     Tidb     `yaml:"tidb"`
	Api      Api      `yaml:"api"`
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Redis    Redis    `yaml:"redis"`
}

// Consumer
type Consumer struct {
	Concurrency int    `yaml:"concurrency"`
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
}

// Api
type Api struct {
	Version int `yaml:"version"`
}

