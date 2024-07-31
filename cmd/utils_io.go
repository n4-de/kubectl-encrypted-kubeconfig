package cmd

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func getPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println() // Print a newline after the password input
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func writeStringToFile(filename string, data string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.WriteString(data)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return err
}
