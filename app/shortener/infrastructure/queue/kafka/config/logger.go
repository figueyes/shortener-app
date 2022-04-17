package config

type Logger interface {
	Printf(string, ...interface{})
}
