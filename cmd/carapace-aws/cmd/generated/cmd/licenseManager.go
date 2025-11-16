package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerCmd = &cobra.Command{
	Use:   "license-manager",
	Short: "License Manager makes it easier to manage licenses from software vendors across multiple Amazon Web Services accounts and on-premises servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerCmd).Standalone()

	})
	rootCmd.AddCommand(licenseManagerCmd)
}
