package log

type Config struct {
	Level string
	File  struct {
		Enabled bool
		File    string
	}
}
