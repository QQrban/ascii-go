package utils

import (
	"os"
	"testing"
)

func TestPrepareOutputFile(t *testing.T) {
	tests := []struct {
		name       string
		flagArg    string
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid TXT Extension",
			flagArg: "--output=testfile.txt",
			wantErr: false,
		},
		{
			name:       "Invalid Extension",
			flagArg:    "--output=testfile.jpeg",
			wantErr:    true,
			errMessage: "OOPS! File extension must be .txt. (i.e. test.txt)  :)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, err := prepareOutputFile(tt.flagArg)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareOutputFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMessage {
				t.Errorf("prepareOutputFile() error message = %v, expected %v", err.Error(), tt.errMessage)
			}
			if gotFile != nil {
				gotFile.Close()
				os.Remove(gotFile.Name())
			}
		})
	}
}

func TestProcessArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantStart int
		wantEnd   int
	}{
		{
			name:      "Flagged and Matched Style",
			args:      []string{"PATH", "--output=test.txt", "Hello", "World", "shadow"},
			wantStart: 2,
			wantEnd:   4,
		},
		{
			name:      "Not Flagged and Matched Style",
			args:      []string{"PATH", "Hello", "World", "shadow"},
			wantStart: 1,
			wantEnd:   3,
		},
		{
			name:      "Flagged and Not Matched Style",
			args:      []string{"PATH", "--output=test.txt", "Hello", "World", "customStyle"},
			wantStart: 2,
			wantEnd:   5,
		},
		{
			name:      "Not Flagged and Not Matched Style",
			args:      []string{"PATH", "Hello", "World", "customStyle"},
			wantStart: 1,
			wantEnd:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd := processArgs(tt.args)
			if gotStart != tt.wantStart || gotEnd != tt.wantEnd {
				t.Errorf("processArgs() = (%v, %v), want (%v, %v)", gotStart, gotEnd, tt.wantStart, tt.wantEnd)
			}
		})
	}
}
