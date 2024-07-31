/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/gopasspw/gopass/pkg/pwgen"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Encrypts a pre-existing KUBECONFIG file",
	Long: `Convert is able to convert a pre-existing KUBECONFIG file into an encrypted KUBECONFIG file.
The convert command also supports in place rewrites.
The automatically generated password can either be written to a file or directly to the systems keyring.
`,
	Run: convert,
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringP("context", "c", "", "context to convert")
	convertCmd.Flags().BoolP("write-config", "w", false, "write config inplace")
	convertCmd.Flags().StringP("outfile", "o", "", "write encrypted config to file")
	convertCmd.Flags().StringP("password-outfile", "p", "", "write generated password unencrypted to file")
	convertCmd.Flags().BoolP("use-keyring", "k", false, "write password directly to keyring")
}

func convert(cmd *cobra.Command, args []string) {
	contextFlag, _ = cmd.Flags().GetString("context")
	writeConfigFlag, _ = cmd.Flags().GetBool("write-config")
	outfileFlag, _ = cmd.Flags().GetString("outfile")
	passwordOutfileFlag, _ = cmd.Flags().GetString("password-outfile")
	useKeyringFlag, _ = cmd.Flags().GetBool("use-keyring")

	config = loadKubeconfig(kubeconfigFlag)

	if len(contextFlag) > 0 {
		contextName := contextFlag
		context := getContext(contextName)
		fmt.Printf("Processing single context: %s\n", contextName)
		convertContext(context, contextName)
	} else {
		fmt.Printf("Iterating over all contexts\n")
		for contextName, context := range config.Contexts {
			convertContext(context, contextName)
		}
	}
}

func convertContext(context *api.Context, contextName string) {
	authInfoName := context.AuthInfo
	clusterName := context.Cluster
	cluster := getCluster(clusterName)
	server := cluster.Server

	fmt.Printf("Context: %s, User: %s, Cluster: %s, Server: %s\n",
		contextName,
		authInfoName,
		clusterName,
		server,
	)

	authInfo := getAuthInfo(authInfoName)

	updated := false
	// only convert Token AuthInfos
	if authInfo.Token != "" {
		updated = true
		password := pwgen.GeneratePassword(32, false)

		// output to password outfile
		if passwordOutfileFlag != "" {
			writeStringToFile(passwordOutfileFlag, password)
			fmt.Printf("Encryption password successfully saved to file '%s'.\n", passwordOutfileFlag)
		}

		// output to keyring
		if useKeyringFlag {
			keyringSet(clusterName, authInfoName, password)
			fmt.Printf("Encryption password successfully saved to keyring (service name: '%s', user name: '%s').\n", getServiceName(clusterName), authInfoName)
		}

		// only output to stderr of really needed
		if !useKeyringFlag && passwordOutfileFlag == "" {
			fmt.Fprintf(os.Stderr, "Encryption password: %s\n", password)
		}

		encryptedData := encryptSymmetric(authInfo.Token, password)

		authInfo.Exec = &api.ExecConfig{
			Command: "kubectl",
			Args: []string{
				"encrypted-kubeconfig",
				"auth",
				"--clustername",
				clusterName,
				"--username",
				authInfoName,
			},
			APIVersion:         "client.authentication.k8s.io/v1",
			ProvideClusterInfo: false,
			InteractiveMode:    api.IfAvailableExecInteractiveMode,
			Env: []api.ExecEnvVar{
				{
					Name:  encryptedTokenEnvVar,
					Value: encryptedData,
				},
			},
		}
		authInfo.Token = "" // remove unencrypted token from config
	}

	if updated {
		if !writeConfigFlag && outfileFlag == "" {
			x, _ := clientcmd.Write(*config)
			fmt.Printf("Output of config:\n%v\n---\n", string(x))
		} else {
			outfile := kubeconfigFlag
			if outfileFlag != "" {
				outfile = outfileFlag
			}
			err := clientcmd.WriteToFile(*config, outfile)
			if err != nil {
				fmt.Printf("Fehler %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("KUBECONFIG successfully saved to file '%s'.\n", outfile)
		}
	} else {
		fmt.Printf("Unable to convert context '%s' (cluster: '%s', user: '%s'). Is this context already converted? ABORTING.\n", contextName, clusterName, authInfoName)
	}
}
