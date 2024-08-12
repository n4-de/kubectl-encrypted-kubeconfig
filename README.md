# kubectl-encrypted-kubeconfig

`kubectl-encrypted-kubeconfig` is a plugin that facilitates authentication to the Kubernetes API by decrypting the access token within an encrypted KUBECONFIG file. This decryption is achieved using the system keyring, ensuring secure and convenient access management.

## Features

- **Secure Authentication:** Decrypt access tokens on the fly using the system keyring for seamless authentication against a Kubernetes cluster.
- **Encryption Conversion:** Easily encrypt existing KUBECONFIG files to enhance security.

## Installation

### Manual installation using krew

1. Download encrypted-kubeconfig.yaml and the os specific package

2. Run the following in the shell

```bash
kubectl krew install --manifest=encrypted-kubeconfig.yaml --archive=encrypted-kubeconfig-[darwin-arm64.tar.gz|linux-amd64.tar.gz|windows-amd64.zip]
```

## Usage
```bash
kubectl-encrypted-kubeconfig [command]
```

### Available Commands
- **auth**: Perform actual authentication against a cluster.
- **completion**: Generate the autocompletion script for the specified shell.
- **convert**: Encrypt a pre-existing KUBECONFIG file.
- **help**: Display help information about any command.
- **keyring**: Execute keyring manipulation functions.

### Global Flags
```
-h, --help: Show help for kubectl-encrypted-kubeconfig.

    --kubeconfig string   Path to KUBECONFIG file
```

### Examples for using the conversion feature

#### Generate an encrypted version of a kubeconfig.
```bash
export KUBECONFIG=/[PATH_TO_KUBECONFIG]/config
kubectl-encrypted-kubeconfig convert -o OUTPUT_FILEPATH -k
```

#### Generate an encrypted version of a kubeconfig and replace the source kubeconfig.
```bash
export KUBECONFIG=/[PATH_TO_KUBECONFIG]/config
kubectl-encrypted-kubeconfig convert -w -k
```

## License
Licensed under Apache v2.0 see [License](https://github.com/n4-de/kubectl-encrypted-kubeconfig/blob/main/LICENSE) for details.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue to discuss improvements or feature requests.
