package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nbutton23/zxcvbn-go"
)

func main() {
	for {
		fmt.Println("Enter password or Ctrl-c to exit:")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')

		passwordStrength := zxcvbn.PasswordStrength(password, nil)

		fmt.Printf("Password score    (0-4): %d\nEstimated entropy (bit): %f\nEstimated time to crack: %s\n\n",
			passwordStrength.Score,
			passwordStrength.Entropy,
			passwordStrength.CrackTimeDisplay,
		)
	}
}
