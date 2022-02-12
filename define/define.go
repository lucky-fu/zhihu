package define

// 多少啊s
type ConfigRedis struct {
	Host     string `mapstructure:"host,omitempty"`
	Port     string `mapstructure:"port,omitempty"`
	DB       int    `mapstructure:"db,omitempty"`
	Password string `mapstructure:"password,omitempty"`

	ConnectTimeout float32 `mapstructure:"connect_timeout,omitempty"`
	WriteTimeout   float32 `mapstructure:"write_timeout,omitempty"`
	ReadTimeout    float32 `mapstructure:"read_timeout,omitempty"`
}
