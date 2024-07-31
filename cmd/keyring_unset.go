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

// unsetCmd represents the unset command
var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove decryption password from keyring",
	Long:  `Use unset to remove a decryption password from your keyring.`,
	Run:   runKeyringUnset,
}

func init() {
	keyringCmd.AddCommand(unsetCmd)
}

func runKeyringUnset(cmd *cobra.Command, args []string) {
	config = loadKubeconfig(kubeconfigFlag)

	contextFlag, _ = cmd.Flags().GetString("context")
	context := getContext(contextFlag)
	authInfoName := context.AuthInfo
	clusterName := context.Cluster

	if authInfoName != "" && clusterName != "" {
		err := keyringUnset(clusterName, authInfoName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Removing password to keyring failed: %v\nCluster: '%s'\nUser: '%s'\n", err, clusterName, authInfoName)
			os.Exit(1)
		}
		fmt.Printf("Password for cluster '%s' and user '%s' was successfully deleted from keyring.\n", clusterName, authInfoName)
	} else {
		fmt.Fprintf(os.Stderr, "Cluster, user or entered password were empty. ABORTING.\n")
		os.Exit(1)
	}

}
