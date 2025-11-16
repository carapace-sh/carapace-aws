package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_setSecurityTokenServicePreferencesCmd = &cobra.Command{
	Use:   "set-security-token-service-preferences",
	Short: "Sets the specified version of the global endpoint token as the token version used for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_setSecurityTokenServicePreferencesCmd).Standalone()

	iam_setSecurityTokenServicePreferencesCmd.Flags().String("global-endpoint-token-version", "", "The version of the global endpoint token.")
	iam_setSecurityTokenServicePreferencesCmd.MarkFlagRequired("global-endpoint-token-version")
	iamCmd.AddCommand(iam_setSecurityTokenServicePreferencesCmd)
}
