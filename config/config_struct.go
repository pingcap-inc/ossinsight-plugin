package config

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Interval
type Interval struct {
	LatestDuring int    `yaml:"latestDuring"`
	Retry        int    `yaml:"retry"`
	RetryWait    int    `yaml:"retryWait"`
	Daily        string `yaml:"daily"`
	Latest       string `yaml:"latest"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Loop
type Loop struct {
	Timeout int `yaml:"timeout"`
	Break   int `yaml:"break"`
}

// Sql
type Sql struct {
	EventsDaily string `yaml:"eventsDaily"`
	Yearly      string `yaml:"yearly"`
}

// Producer
type Producer struct {
	Topic string `yaml:"topic"`
	Retry int    `yaml:"retry"`
}

// Server
type Server struct {
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
	Port      int    `yaml:"port"`
}

// Pulsar
type Pulsar struct {
	Env      string   `yaml:"env"`
	DevHost  string   `yaml:"devHost"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Lark
type Lark struct {
	Webhook             string `yaml:"webhook"`
	SignKey             string `yaml:"signKey"`
	MinimumBreak        int    `yaml:"minimumBreak"`
	ErrorTolerance      int    `yaml:"errorTolerance"`
	ErrorToleranceBreak int    `yaml:"errorToleranceBreak"`
}

// Risingwave
type Risingwave struct {
	Db       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Api
type Api struct {
	Version int `yaml:"version"`
}

// Redis
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Lua      Lua    `yaml:"lua"`
}

// Lua
type Lua struct {
	MergeLatest string `yaml:"mergeLatest"`
}

// Config
type Config struct {
	Github     Github     `yaml:"github"`
	Tidb       Tidb       `yaml:"tidb"`
	Risingwave Risingwave `yaml:"risingwave"`
	Api        Api        `yaml:"api"`
	Server     Server     `yaml:"server"`
	Redis      Redis      `yaml:"redis"`
	Pulsar     Pulsar     `yaml:"pulsar"`
	Interval   Interval   `yaml:"interval"`
	Log        Log        `yaml:"log"`
	Lark       Lark       `yaml:"lark"`
}

// Github
type Github struct {
	Loop   Loop     `yaml:"loop"`
	Tokens []string `yaml:"tokens"`
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

