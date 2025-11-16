package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystoreCmd = &cobra.Command{
	Use:   "identitystore",
	Short: "The Identity Store service used by IAM Identity Center provides a single place to retrieve all of your identities (users and groups).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(identitystoreCmd).Standalone()

	})
	rootCmd.AddCommand(identitystoreCmd)
}
