/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-encrypted-kubeconfig",
	Short: "Authentication plugin for encrypted KUBECONFIG files.",
	Long:  `This plugin allows to authenticate to Kubernetes API by decrypting the accesstoken inside of an encrypted KUBECONFIG. This is done using the system keyring.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	kubeconfigFallback := os.Getenv("KUBECONFIG")
	if kubeconfigFallback == "" {
		kubeconfigFallback = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	rootCmd.PersistentFlags().StringVar(&kubeconfigFlag, "kubeconfig", kubeconfigFallback, "Path to KUBECONFIG file")
}
