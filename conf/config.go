package conf

type Config struct {
	MySQL   MySQLConfig   `mapstructure:"mysql"`
	Redis   RedisConfig   `mapstructure:"redis"`
	Timeout TimeoutConfig `mapstructure:"timeout"`
	Server  ServerConfig  `mapstructure:"server"`
	Jwt     JwtConfig     `mapstructure:"jwt"`
	Minio   MinioConfig   `mapstructure:"minio"`
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
	Port        int    `mapstructure:"port"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"DB"`
	PoolSize    int    `mapstructure:"poolSize"`
	MinIdleConn int    `mapstructure:"minIdleConn"`
}

type TimeoutConfig struct {
	DelayHeartbeat   int `mapstructure:"delayHeartbeat"`
	HeartbeatHz      int `mapstructure:"heartbeatHz"`
	HeartbeatMaxTime int `mapstructure:"heartbeatMaxTime"`
	RedisOnlineTime  int `mapstructure:"redisOnlineTime"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
	UDP  int `mapstructure:"udp"`
}

type JwtConfig struct {
	SecretKey string `mapstructure:"secretKey"`
}

type MinioConfig struct {
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	Bucket    string `mapstructure:"bucket"`
	Endpoint  string `mapstructure:"endpoint"`
	UseSSL    bool   `mapstructure:"useSSL"`
}
