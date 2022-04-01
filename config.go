package themisclient

import (
	"go.themis.run/themisclient/loadbalance"
)

type Config struct {
	ServerName    string
	ServerAddress string

	LoadBalancerName string

	RetryNum int
}

func NewConfigration(opts ...Option) *Config {
	config := DefaultConfigration()

	for _, o := range opts {
		o(config)
	}

	return config
}

var DefaultConfigration = func() *Config {
	return &Config{
		LoadBalancerName: loadbalance.DefaultName,
		RetryNum:         3,
	}
}

type Option func(*Config)

func WithServerName(name string) Option {
	return func(c *Config) {
		c.ServerName = name
	}
}

func WithServerAddress(addr string) Option {
	return func(c *Config) {
		c.ServerAddress = addr
	}
}

func WithLoadBalancerName(name string) Option {
	return func(c *Config) {
		c.LoadBalancerName = name
	}
}

func WithRetryNum(num int) Option {
	return func(c *Config) {
		c.RetryNum = num
	}
}
