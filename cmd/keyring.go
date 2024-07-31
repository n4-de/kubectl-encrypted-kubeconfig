/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// keyringCmd represents the keyring command
var keyringCmd = &cobra.Command{
	Use:   "keyring",
	Short: "Keyring manipulation functions",
	Long:  `Use the keyring command to interact with your keyring.`,
}

func init() {
	rootCmd.AddCommand(keyringCmd)

	keyringCmd.PersistentFlags().StringP("context", "c", "", "context to specify password for")
}

func getServiceName(service string) string {
	return serviceNamePrefix + service + serviceNameSuffix
}

func keyringSet(service string, username string, password string) error {
	return keyring.Set(getServiceName(service), username, password)
}

func keyringGet(service string, username string) string {
	password, err := keyring.Get(getServiceName(service), username)
	if err != nil {
		return ""
	}

	return password
}

func keyringUnset(service string, username string) error {
	return keyring.Delete(getServiceName(service), username)
}
