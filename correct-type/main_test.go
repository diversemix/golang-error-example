package main

import "testing"

func TestReadCSV(t *testing.T) {
	tests := map[string]struct {
		filename      string
		errorExpected bool
	}{
		"loads a good CSV":         {filename: "good.csv", errorExpected: false},
		"errors on bad CSV":        {filename: "bad.csv", errorExpected: true},
		"errors on file not found": {filename: "foobar.csv", errorExpected: true},
	}

	for name, context := range tests {
		result, err := ReadCSV(context.filename)
		if context.errorExpected && err == nil {
			t.Errorf("%s: Expected an error for file %s, received none.", name, context.filename)
		}
		if !context.errorExpected {
			if err != nil {
				t.Errorf("%s: Received an error %v when not expecting one", name, err)
			}
			if result == "" {
				t.Errorf("%s: Expecting a return value got a blank string", name)
			}
		}
	}

}
