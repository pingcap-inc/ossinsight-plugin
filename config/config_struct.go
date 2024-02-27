package config

// Tidb
type Tidb struct {
	Db       string `yaml:"db"`
	Sql      Sql    `yaml:"sql"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Disable
type Disable struct {
	Interval bool `yaml:"interval"`
	Producer bool `yaml:"producer"`
}

// Pulsar
type Pulsar struct {
	Env      string   `yaml:"env"`
	DevHost  string   `yaml:"devHost"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

// Consumer
type Consumer struct {
	Topic       string `yaml:"topic"`
	Name        string `yaml:"name"`
	Concurrency int    `yaml:"concurrency"`
}

// Risingwave
type Risingwave struct {
	Db       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Config
type Config struct {
	Server     Server     `yaml:"server"`
	Log        Log        `yaml:"log"`
	Lark       Lark       `yaml:"lark"`
	Redis      Redis      `yaml:"redis"`
	Github     Github     `yaml:"github"`
	Interval   Interval   `yaml:"interval"`
	Tidb       Tidb       `yaml:"tidb"`
	Api        Api        `yaml:"api"`
	Disable    Disable    `yaml:"disable"`
	Pulsar     Pulsar     `yaml:"pulsar"`
	Risingwave Risingwave `yaml:"risingwave"`
}

// Log
type Log struct {
	Level  string `yaml:"level"`
	File   string `yaml:"file"`
	Format string `yaml:"format"`
}

// Lua
type Lua struct {
	MergeLatest string `yaml:"mergeLatest"`
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

// Redis
type Redis struct {
	Db       int    `yaml:"db"`
	Lua      Lua    `yaml:"lua"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

// Interval
type Interval struct {
	RetryWait    int    `yaml:"retryWait"`
	Daily        string `yaml:"daily"`
	Latest       string `yaml:"latest"`
	LatestDuring int    `yaml:"latestDuring"`
	Retry        int    `yaml:"retry"`
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
	Port      int    `yaml:"port"`
	Health    string `yaml:"health"`
	SyncEvent string `yaml:"syncEvent"`
}

// Lark
type Lark struct {
	Webhook             string `yaml:"webhook"`
	SignKey             string `yaml:"signKey"`
	MinimumBreak        int    `yaml:"minimumBreak"`
	ErrorTolerance      int    `yaml:"errorTolerance"`
	ErrorToleranceBreak int    `yaml:"errorToleranceBreak"`
}

// Api
type Api struct {
	Version int `yaml:"version"`
}

