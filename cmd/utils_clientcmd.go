package cmd

import (
	"fmt"
	"os"

	"golang.org/x/exp/maps"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func loadKubeconfig(kubeconfigFilename string) *api.Config {
	config, err := clientcmd.LoadFromFile(kubeconfigFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading KUBECONFIG file: %v\nFilename: '%s'\n", err, kubeconfigFilename)
		os.Exit(1)
	}

	return config
}

func getContext(contextName string) *api.Context {
	var context *api.Context

	if contextName != "" {
		// contextFlag is set via command line
		context = config.Contexts[contextName]
	} else {
		// output a list if more than 1 context was found
		if len(config.Contexts) > 1 {
			fmt.Printf("Found multiple contexts - please specify one.\n")
			fmt.Printf("Found contexts:\n")
			for contextNameKey, context := range config.Contexts {
				fmt.Printf("  - name: %s\n", contextNameKey)
				fmt.Printf("    cluster: %s\n", context.Cluster)
				fmt.Printf("    user: %s\n", context.AuthInfo)
			}
			os.Exit(0)
		} else {
			// use the one and only available context as default
			keys := maps.Keys(config.Contexts)
			context = config.Contexts[keys[0]]
		}
	}

	if context == nil {
		fmt.Fprintf(os.Stderr, "Context with name '%s' not found. ABORTING.\n", contextName)
		os.Exit(1)
	}

	return context
}

func getAuthInfo(authInfoName string) *api.AuthInfo {
	authInfo := config.AuthInfos[authInfoName]
	if authInfo == nil {
		fmt.Printf("ERROR: auth info for reference \"%s\" not found in current KUBECONFIG\n", authInfoName)
		os.Exit(1)
	}

	return authInfo
}

func getCluster(clusterName string) *api.Cluster {
	cluster := config.Clusters[clusterName]
	if cluster == nil {
		fmt.Printf("ERROR: cluster name for reference \"%s\" not found in current KUBECONFIG\n", clusterName)
		os.Exit(1)
	}

	return cluster
}
