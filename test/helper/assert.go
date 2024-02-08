package helper

// These functions may be transferred to a separate lib from hashdog

import (
	"errors"
	"testing"
)

type Adapter struct{ t *testing.T }

func NewAdapter(t *testing.T) Adapter {
	t.Helper()

	return Adapter{t: t}
}

// Assert tests if two values are different, print an error if they are.
func (suite Adapter) Assert(got, expected any) {
	suite.t.Helper()
	if got != expected {
		suite.t.Errorf("Got: %v | Expected: %v\n", got, expected)
	}
}

// Assertf tests if two values are different, fail tests if they are.
func (suite Adapter) Assertf(got, expected any) {
	suite.t.Helper()
	if got != expected {
		suite.t.Fatalf("Got: %v | Expected: %v\n", got, expected)
	}
}

// AssertErrIs tests if two errors are different, print an error if they are.
func (suite Adapter) AssertErrIs(got, expected error) {
	suite.t.Helper()
	if !errors.Is(got, expected) {
		suite.t.Errorf("Got: %v | Expected: %v\n", got, expected)
	}
}

// AssertErrIsf tests if two errors are different, fail tests if they are.
func (suite Adapter) AssertErrIsf(got, expected error) {
	suite.t.Helper()
	if !errors.Is(got, expected) {
		suite.t.Fatalf("Got: %v | Expected: %v\n", got, expected)
	}
}

// AssertErrAs tests if two errors are of different types, print an error if they are.
func (suite Adapter) AssertErrAs(got error, expected any) {
	suite.t.Helper()
	if !errors.As(got, &expected) {
		suite.t.Errorf("Got: %v | Expected: %v\n", got, expected)
	}
}

// AssertErrAsf tests if two errors are of different types, fail tests if they are.
func (suite Adapter) AssertErrAsf(got error, expected any) {
	suite.t.Helper()
	if !errors.As(got, &expected) {
		suite.t.Fatalf("Got: %v | Expected: %v\n", got, expected)
	}
}
