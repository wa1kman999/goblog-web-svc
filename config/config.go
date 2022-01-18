package config

// Config 全局配置
type Config struct {
	System *System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    *JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Consul *Consul `mapstructure:"consul" json:"consul" yaml:"consul"`
}
