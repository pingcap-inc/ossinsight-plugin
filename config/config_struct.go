package config

// Redis
type Redis struct {
	Lua      Lua    `yaml:"lua"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// Producer
type Producer struct {
	Topic string `yaml:"topic"`
	Retry int    `yaml:"retry"`
}

// Sql
type Sql struct {
	EventsDaily string `yaml:"eventsDaily"`
	Yearly      string `yaml:"yearly"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Server
type Server struct {
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
	Port      int    `yaml:"port"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Config
type Config struct {
	Redis    Redis    `yaml:"redis"`
	Pulsar   Pulsar   `yaml:"pulsar"`
	Github   Github   `yaml:"github"`
	Tidb     Tidb     `yaml:"tidb"`
	Api      Api      `yaml:"api"`
	Disable  Disable  `yaml:"disable"`
	Server   Server   `yaml:"server"`
	Log      Log      `yaml:"log"`
	Interval Interval `yaml:"interval"`
}

// Lua
type Lua struct {
	MergeLatest string `yaml:"mergeLatest"`
}

// Pulsar
type Pulsar struct {
	Consumer Consumer `yaml:"consumer"`
	Env      string   `yaml:"env"`
	DevHost  string   `yaml:"devHost"`
	Host     string   `yaml:"host"`
	Audience string   `yaml:"audience"`
	Keypath  string   `yaml:"keypath"`
	Producer Producer `yaml:"producer"`
}

// Consumer
type Consumer struct {
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
	Topic       string `yaml:"topic"`
}

// Github
type Github struct {
	Loop   Loop     `yaml:"loop"`
	Tokens []string `yaml:"tokens"`
}

// Api
type Api struct {
	Version int `yaml:"version"`
}

// Interval
type Interval struct {
	Daily        string `yaml:"daily"`
	Latest       string `yaml:"latest"`
	LatestDuring int    `yaml:"latestDuring"`
	Retry        int    `yaml:"retry"`
	RetryWait    int    `yaml:"retryWait"`
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

// Disable
type Disable struct {
	Producer bool `yaml:"producer"`
	Interval bool `yaml:"interval"`
}
