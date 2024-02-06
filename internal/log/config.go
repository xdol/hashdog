package log

type Config struct {
	level string
	file  struct {
		enabled bool
		file    string
	}
}
