package main

import (
	"os"
	"testing"
)

func TestReadCSV(t *testing.T) {
	tests := map[string]struct {
		filename      string
		expectedError func(err error) bool
	}{
		"loads a good CSV":         {filename: "good.csv", expectedError: noError},
		"errors on bad CSV":        {filename: "bad.csv", expectedError: IsInvalidCSVContent},
		"errors on file not found": {filename: "foobar.csv", expectedError: os.IsNotExist},
	}

	for name, context := range tests {
		_, err := ReadCSV(context.filename)
		if !context.expectedError(err) {
			t.Errorf("%s: Got unexpected error or error of unexpected type for file %s.", name, context.filename)
		}
	}

}

func noError(err error) bool {
	return err == nil
}
