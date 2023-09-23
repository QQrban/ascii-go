package utils

import "testing"

func TestGetStyle(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		expected string
	}{
		{"With style shadow", "Do you want some magic? shadow", "shadow"},
		{"With style standard", " -output=hello.txt standard", "standard"},
		{"With style thinkertoy", "I love Codin'!!! thinkertoy", "thinkertoy"},
		{"Without style", "I hate pineapple pizza ", "standard"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStyle(tt.args); got != tt.expected {
				t.Errorf("HandleArgs() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
