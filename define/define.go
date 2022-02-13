package define

// ConfigRedis ...
type ConfigRedis struct {
	Host     string `mapstructure:"host,omitempty"`
	Port     string `mapstructure:"port,omitempty"`
	DB       int    `mapstructure:"db,omitempty"`
	Password string `mapstructure:"password,omitempty"`

	ConnectTimeout float32 `mapstructure:"connect_timeout,omitempty"`
	WriteTimeout   float32 `mapstructure:"write_timeout,omitempty"`
	ReadTimeout    float32 `mapstructure:"read_timeout,omitempty"`
}

// ConfigDB ...
type ConfigDB struct {
	Driver        string `mapstructure:"driver,omitempty"`
	WriteHost     string `mapstructure:"write_host,omitempty"`
	WriteUsername string `mapstructure:"write_username,omitempty"`
	WritePassword string `mapstructure:"write_password,omitempty"`
	ReadHost      string `mapstructure:"read_host,omitempty"`
	ReadUsername  string `mapstructure:"read_username,omitempty"`
	ReadPassword  string `mapstructure:"read_password,omitempty"`
	Database      string `mapstructure:"database,omitempty"`
	Charset       string `mapstructure:"charset,omitempty"`
	Collation     string `mapstructure:"collation,omitempty"`
}
