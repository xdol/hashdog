package config

import (
	"os"

	"git.xdol.org/xdol/go-yaml"
)

type LoadYamlError struct {
	path string
	err  error
}

func (m *LoadYamlError) Error() string {
	return "Cannot load " + m.path + ": " + m.err.Error()
}

type VerifyConfigError struct {
	problem string
}

func (m *VerifyConfigError) Error() string {
	return "Error in config: " + m.problem
}

func LoadYaml(file string) (Config, error) {
	var config Config
	source, err := os.ReadFile(file)
	if err != nil {
		return config, &LoadYamlError{path: file, err: err}
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return config, &LoadYamlError{path: file, err: err}
	}

	return config, nil
}

func VerifyConfig(config Config) error { //nolint:cyclop
	if !config.Bruteforce.Charset.Enabled && !config.Bruteforce.Wordlist.Enabled && !config.Bruteforce.Rainbow.Enabled {
		return &VerifyConfigError{problem: "No bruteforce mode enabled."}
	}
	if config.Bruteforce.Charset.Enabled && config.Bruteforce.Rainbow.Enabled {
		return &VerifyConfigError{problem: "Multiple bruteforce mode enabled."}
	}
	if config.Bruteforce.Charset.Enabled && config.Bruteforce.Wordlist.Enabled {
		return &VerifyConfigError{problem: "Multiple bruteforce mode enabled."}
	}
	if config.Bruteforce.Rainbow.Enabled && config.Bruteforce.Wordlist.Enabled {
		return &VerifyConfigError{problem: "Multiple bruteforce mode enabled."}
	}

	allowedValue := []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	found := false
	for _, v := range allowedValue {
		if v == config.Log.Level {
			found = true
		}
	}
	if !found {
		return &VerifyConfigError{problem: "Log level is not valid."}
	}

	return nil
}
