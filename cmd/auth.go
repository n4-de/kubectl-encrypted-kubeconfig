/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Actual authentication against a cluster",
	Long: `The actual authentication against a cluster is done using this command.
You need to specify a cluster name and also a user name.	
This command works as a client-go credentials plugin.
It uses the API version 'client.authentication.k8s.io/v1'.`,
	Run: auth,
}

func auth(cmd *cobra.Command, args []string) {
	clusternameFlag, _ = cmd.Flags().GetString("clustername")
	usernameFlag, _ = cmd.Flags().GetString("username")
	encryptedToken := os.Getenv(encryptedTokenEnvVar)

	if clusternameFlag == "" || usernameFlag == "" || encryptedToken == "" {
		fmt.Fprintf(os.Stderr, "clustername, username or encrypted token not set. ABORTING.\n")
		os.Exit(1)
	}

	password := keyringGet(clusternameFlag, usernameFlag)
	if password == "" {
		fmt.Fprintf(os.Stderr, "Password not found in keyring. Please set it with keyring set. ABORTING.\n")
		os.Exit(1)
	}

	token := decryptSymmetric(encryptedToken, password)
	if token == "" {
		fmt.Fprintf(os.Stderr, "Decryption of token failed. Did you specify the correct password? ABORTING.\n")
		os.Exit(1)
	}

	fmt.Printf(`{
		"apiVersion": "client.authentication.k8s.io/v1",
		"kind": "ExecCredential",
		"status": {
		  "token": "%s"
		}
	  }
	  `, token)
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().String("username", "", "user name used to authenticate against the cluster")
	authCmd.Flags().String("clustername", "", "cluster name to authenticate to")
}
