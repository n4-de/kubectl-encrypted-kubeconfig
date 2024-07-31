/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"filippo.io/age"
	"filippo.io/age/armor"
)

func encryptSymmetric(data string, password string) string {
	recipient, _ := age.NewScryptRecipient(password)

	return encryptData(data, recipient)
}

func decryptSymmetric(data string, password string) string {
	identity, _ := age.NewScryptIdentity(password)

	return decryptData(data, identity)
}

func encryptData(data string, recipient age.Recipient) string {
	var encrypted bytes.Buffer

	armor := armor.NewWriter(&encrypted)
	writer, err := age.Encrypt(armor, recipient)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	_, err = io.WriteString(writer, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	err = writer.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	err = armor.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	return encrypted.String()
}

func decryptData(data string, identity age.Identity) string {
	var decrypted bytes.Buffer

	stringReader := strings.NewReader(data)
	armor := armor.NewReader(stringReader)

	reader, err := age.Decrypt(armor, identity)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR while decrypting: %v\n", err)
		return ""
	}

	if _, err := io.Copy(&decrypted, reader); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR while copying: %v\n", err)
		return ""
	}

	return decrypted.String()
}
