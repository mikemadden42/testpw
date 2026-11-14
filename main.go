package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ccojocar/zxcvbn-go"
	"golang.org/x/term"
)

func main() {
	fd := int(os.Stdin.Fd()) //nolint:gosec // file descriptors fit in an int
	interactive := term.IsTerminal(fd)
	reader := bufio.NewReader(os.Stdin)

	for {
		if interactive {
			fmt.Println("Enter password or Ctrl-c to exit:")
		}

		password, err := readPassword(fd, interactive, reader)
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading password:", err)
			os.Exit(1)
		}

		fmt.Println(report(password))
	}
}

// readPassword reads one password without echo from a terminal, or one line
// from piped input. It returns io.EOF once piped input is exhausted.
func readPassword(fd int, interactive bool, reader *bufio.Reader) (string, error) {
	if interactive {
		b, err := term.ReadPassword(fd)
		fmt.Println()
		return string(b), err
	}

	line, err := reader.ReadString('\n')
	if errors.Is(err, io.EOF) && line != "" {
		err = nil
	}
	return strings.TrimRight(line, "\r\n"), err
}

// report scores password and formats the result.
func report(password string) string {
	s := zxcvbn.PasswordStrength(password, nil)
	return fmt.Sprintf("Password score    (0-4): %d\nEstimated entropy (bit): %.2f\nEstimated time to crack: %s\n",
		s.Score,
		s.Entropy,
		s.CrackTimeDisplay,
	)
}
