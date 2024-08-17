package config

type Config struct {
	RunAddress           string `valid:"minstringlength(3)"`
	LogLevel             string `valid:"in(debug|info|warn|error|fatal)"`
	DatabaseURI          string
	AccrualSystemAddress string
	AuthKey              string
	AuthKeyExpired       int
}
