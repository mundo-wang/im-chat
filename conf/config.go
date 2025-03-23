package conf

type Config struct {
	MySQL   MySQLConfig   `mapstructure:"mysql"`
	Redis   RedisConfig   `mapstructure:"redis"`
	OSS     OSSConfig     `mapstructure:"oss"`
	Timeout TimeoutConfig `mapstructure:"timeout"`
	Port    PortConfig    `mapstructure:"port"`
}

type MySQLConfig struct {
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
	IP       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	IP          string `mapstructure:"ip"`
	Port        string `mapstructure:"port"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"DB"`
	PoolSize    int    `mapstructure:"poolSize"`
	MinIdleConn int    `mapstructure:"minIdleConn"`
}

type OSSConfig struct {
	Endpoint        string `mapstructure:"Endpoint"`
	AccessKeyId     string `mapstructure:"AccessKeyId"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Bucket          string `mapstructure:"Bucket"`
}

type TimeoutConfig struct {
	DelayHeartbeat   int `mapstructure:"DelayHeartbeat"`
	HeartbeatHz      int `mapstructure:"HeartbeatHz"`
	HeartbeatMaxTime int `mapstructure:"HeartbeatMaxTime"`
	RedisOnlineTime  int `mapstructure:"RedisOnlineTime"`
}

type PortConfig struct {
	Server int `mapstructure:"server"`
	UDP    int `mapstructure:"udp"`
}
