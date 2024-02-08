package config_test

import (
	"testing"

	"git.xdol.org/xdol/hashdog/internal/bruteforce"
	"git.xdol.org/xdol/hashdog/internal/bruteforce/charset"
	"git.xdol.org/xdol/hashdog/internal/bruteforce/rainbow"
	"git.xdol.org/xdol/hashdog/internal/config"
	"git.xdol.org/xdol/hashdog/internal/log"
	"git.xdol.org/xdol/hashdog/test/helper"
)

// Test suite structure.
type yamlTestSuite struct {
	t *testing.T
	a helper.Adapter
}

// TestYaml calls all the tests.
func TestYaml(t *testing.T) {
	// Enable parallelism
	t.Parallel()

	// Initialize the helper's adapter
	assertHelper := helper.NewAdapter(t)

	// Initialize the test suite
	suite := yamlTestSuite{t: t, a: assertHelper}

	// Call the tests
	suite.TestFileInvalid()
	suite.TestFileLoad()
	suite.TestFileNotExist()
	suite.TestVerifyConfigMultipleMode()
	suite.TestLogLevel()
	suite.TestLogLevelInvalid()
}

// TestFileLoad tests if a yaml file loads correctly
// and if the loaded content if correct.
func (suite yamlTestSuite) TestFileLoad() { //nolint:funlen
	// Load the file and get the content
	got, err := config.LoadYaml("resources/test_valid.yml")
	suite.a.AssertErrIs(err, nil)

	// Define what is expected
	expected := config.Config{
		Bruteforce: bruteforce.Config{
			Length:  8,
			Threads: 0,
			Rainbow: rainbow.Config{
				Enabled: false,
				Path:    "example.txt",
			},
			Wordlist: rainbow.Config{
				Enabled: false,
				Path:    "example.txt",
			},
			Charset: charset.Config{
				Enabled: true,
				Default: "alphabetloweruppernum",
				List: map[string]string{
					"alphabetlower":         "abcdefghijklmnopqrstuvwxyz",
					"alphabetlowerupper":    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
					"alphabetloweruppernum": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz012345678",
					"alphabetupper":         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
					"num":                   "0123456789",
				},
			},
		},
		Log: log.Config{
			Level: "ERROR",
			File: struct {
				Enabled bool
				File    string
			}{
				Enabled: false,
				File:    "",
			},
		},
	}

	// Test the values of Config.Bruteforce.Rainbow
	suite.a.Assert(got.Bruteforce.Rainbow, expected.Bruteforce.Rainbow)

	// Test the values of Config.Bruteforce.Wordlist
	suite.a.Assert(got.Bruteforce.Wordlist, expected.Bruteforce.Wordlist)

	// Test the values of Config.Bruteforce.Length
	suite.a.Assert(got.Bruteforce.Length, expected.Bruteforce.Length)

	// Test the values of Config.Bruteforce.Threads
	suite.a.Assert(got.Bruteforce.Threads, expected.Bruteforce.Threads)

	// Test the values of Config.Bruteforce.Charset.Default
	suite.a.Assert(got.Bruteforce.Charset.Default, expected.Bruteforce.Charset.Default)

	// Test the values of Config.Bruteforce.Charset.List
	for key, value := range expected.Bruteforce.Charset.List {
		suite.a.Assert(got.Bruteforce.Charset.List[key], value)
	}

	// Test the values of Config.Log
	suite.a.Assert(got.Log, expected.Log)
}

// TestFileInvalid tests if there isn't an error despite that there should be one.
func (suite yamlTestSuite) TestFileInvalid() {
	_, got := config.LoadYaml("resources/test_invalid.yml")
	suite.a.AssertErrAs(got, &config.LoadYamlError{})
}

// TestFileNotExist tests if there's a valid error when loading a file that does not exist.
func (suite yamlTestSuite) TestFileNotExist() {
	_, got := config.LoadYaml("notexist.yaml")
	suite.a.AssertErrAs(got, &config.LoadYamlError{})
}

// TestVerifyConfigMultipleMode tests if there's a valid error when multiple modes are enabled.
func (suite yamlTestSuite) TestVerifyConfigMultipleMode() {
	fileList := []string{
		"resources/wordlist_charset_rainbow.yaml",
		"resources/wordlist_charset.yaml",
		"resources/wordlist_rainbow.yaml",
		"resources/rainbow_charset.yaml",
	}

	for _, file := range fileList {
		s, err := config.LoadYaml(file)
		suite.a.AssertErrIs(err, nil)

		got := config.VerifyConfig(s)
		suite.a.AssertErrAs(got, &config.VerifyConfigError{})
	}
}

// TestLogLevel tests if there's no errors when loading different configs with valid log levels.
func (suite yamlTestSuite) TestLogLevel() {
	// string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	fileList := []string{
		"resources/log_trace.yaml",
		"resources/log_debug.yaml",
		"resources/log_info.yaml",
		"resources/log_warn.yaml",
		"resources/log_error.yaml",
		"resources/log_fatal.yaml",
	}

	for _, file := range fileList {
		s, err := config.LoadYaml(file)
		suite.a.AssertErrIs(err, nil)

		got := config.VerifyConfig(s)
		suite.a.AssertErrIs(got, nil)
	}
}

// TestLogLevelInvalid tests if there's a valid error when verifying an invalid config.
func (suite yamlTestSuite) TestLogLevelInvalid() {
	s, err := config.LoadYaml("resources/log_invalid.yaml")
	if err != nil {
		suite.t.Fatal(err)
	}

	got := config.VerifyConfig(s)
	suite.a.AssertErrAs(got, &config.VerifyConfigError{})
}
