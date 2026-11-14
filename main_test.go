package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReport(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		wantScore string
	}{
		{"common", "password", "Password score    (0-4): 0\n"},
		{"strong", "correct-horse-battery-staple-91!", "Password score    (0-4): 4\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := report(tt.password)
			if !strings.HasPrefix(got, tt.wantScore) {
				t.Errorf("report(%q) = %q, want prefix %q", tt.password, got, tt.wantScore)
			}
		})
	}
}

func TestReadPasswordPiped(t *testing.T) {
	// Covers newline trimming, CRLF, a final line with no newline, and EOF.
	reader := bufio.NewReader(strings.NewReader("first\nsecond\r\nthird"))
	for _, want := range []string{"first", "second", "third"} {
		got, err := readPassword(0, false, reader)
		if err != nil {
			t.Fatalf("readPassword() error = %v", err)
		}
		if got != want {
			t.Errorf("readPassword() = %q, want %q", got, want)
		}
	}
	if _, err := readPassword(0, false, reader); !errors.Is(err, io.EOF) {
		t.Errorf("readPassword() at end error = %v, want io.EOF", err)
	}
}

func TestReportIgnoresNewline(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("Tr0ub4dor&3\n"))
	got, err := readPassword(0, false, reader)
	if err != nil {
		t.Fatal(err)
	}
	if report(got) != report("Tr0ub4dor&3") {
		t.Error("trailing newline changed the score")
	}
}
