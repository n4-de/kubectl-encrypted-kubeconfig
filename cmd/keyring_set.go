/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Store decryption password in keyring",
	Long: `Use set to store the decryption password in your systems keyring.
While specifing the password the password itself is not validated.

You don't have to specify a context if your KUBECONFIG file only contains a single context.
`,
	Run: runKeyringSet,
}

func runKeyringSet(cmd *cobra.Command, args []string) {
	config = loadKubeconfig(kubeconfigFlag)

	contextFlag, _ = cmd.Flags().GetString("context")
	context := getContext(contextFlag)
	authInfoName := context.AuthInfo
	clusterName := context.Cluster

	fmt.Printf("You are about to specify a password for\n")
	fmt.Printf("  cluster: %s\n  user: %s\n", clusterName, authInfoName)
	password, err := getPassword("Please specify your given encryption password: ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Entering a password failed: %v\n", err)
		os.Exit(1)
	}

	if authInfoName != "" && clusterName != "" && password != "" {
		err = keyringSet(clusterName, authInfoName, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Writing password to keyring failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Password for cluster '%s' and user '%s' was successfully written to keyring.\n", clusterName, authInfoName)
	} else {
		fmt.Fprintf(os.Stderr, "Cluster, user or entered password were empty. ABORTING.\n")
		os.Exit(1)
	}
}

func init() {
	keyringCmd.AddCommand(setCmd)
}
