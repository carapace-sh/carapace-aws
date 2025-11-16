package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Retrieves details about your account's [limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html) and usage in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getAccountSettingsCmd).Standalone()

	lambdaCmd.AddCommand(lambda_getAccountSettingsCmd)
}
