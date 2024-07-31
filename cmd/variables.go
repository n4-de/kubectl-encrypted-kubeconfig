package cmd

import "k8s.io/client-go/tools/clientcmd/api"

var (
	encryptedTokenEnvVar = "KUBE_ENCRYPTED_TOKEN"
	serviceNamePrefix    = "K8S-cluster-"
	serviceNameSuffix    = ""
)

var (
	config *api.Config // contains Kubeconfig
)

var (
	kubeconfigFlag      string // global flag
	contextFlag         string // used in convert and keyring set/unset
	writeConfigFlag     bool   // used in convert
	outfileFlag         string // used in convert
	passwordOutfileFlag string // used in convert
	useKeyringFlag      bool   // used in convert
	usernameFlag        string // used in auth
	clusternameFlag     string // used in auth
)
