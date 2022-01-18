package config

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 端口值
}
