package entity

type Provider struct {
	Enable  bool   `yaml:"enable"`
	Timeout string `yaml:"timeout"`
}
