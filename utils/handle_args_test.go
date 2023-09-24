package utils

// import (
// 	"os"
// 	"testing"
// )

// func TestHandleArgs(t *testing.T) {

// 	originalStdout := os.Stdout
// 	os.Stdout, _ = os.Open(os.DevNull)
// 	defer func() { os.Stdout = originalStdout }()

// 	tests := []struct {
// 		name     string
// 		args     string
// 		expected string
// 	}{
// 		{"With flag and filename", "--output=testfile.txt", "testfile.txt"},
// 		{"With flag but no filename", "--output=", "default.txt"},
// 		{"Without flag", "someOtherArgument", ""},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got, _ := HandleArgs(tt.args); got != tt.expected {
// 				t.Errorf("HandleArgs() = %v, expected %v", got, tt.expected)
// 			}
// 		})
// 	}
// }

// func TestIsFlagged(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		args     string
// 		expected bool
// 	}{
// 		{"With flag and filename", "--output=testfile.txt", true},
// 		{"With flag but no filename", "--output=", true},
// 		{"Without flag", "someOtherArgument", false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := IsFlagged(tt.args); got != tt.expected {
// 				t.Errorf("IsFlagged() = %v, expected %v", got, tt.expected)
// 			}
// 		})
// 	}
// }
