/*
Copyright © 2024 Bernhard J. M. Grün <bernhard.gruen@n4.de>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version_name string
	version      = "dev"
	commit_rev   = "dev"
	build_date   = "n/a"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Long:  `Print the version number and build date.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\nCommit rev: %s\nBuild date: %s\n", version_name, version, commit_rev, build_date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
