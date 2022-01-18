package config

type Consul struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 主机
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 端口值
}
