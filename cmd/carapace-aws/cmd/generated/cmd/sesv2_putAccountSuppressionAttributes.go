package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putAccountSuppressionAttributesCmd = &cobra.Command{
	Use:   "put-account-suppression-attributes",
	Short: "Change the settings for the account-level suppression list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putAccountSuppressionAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putAccountSuppressionAttributesCmd).Standalone()

		sesv2_putAccountSuppressionAttributesCmd.Flags().String("suppressed-reasons", "", "A list that contains the reasons that email addresses will be automatically added to the suppression list for your account.")
	})
	sesv2Cmd.AddCommand(sesv2_putAccountSuppressionAttributesCmd)
}
