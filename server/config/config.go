package config

type Server struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
}
