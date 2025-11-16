package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getSuppressedDestinationCmd = &cobra.Command{
	Use:   "get-suppressed-destination",
	Short: "Retrieves information about a specific email address that's on the suppression list for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getSuppressedDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getSuppressedDestinationCmd).Standalone()

		sesv2_getSuppressedDestinationCmd.Flags().String("email-address", "", "The email address that's on the account suppression list.")
		sesv2_getSuppressedDestinationCmd.MarkFlagRequired("email-address")
	})
	sesv2Cmd.AddCommand(sesv2_getSuppressedDestinationCmd)
}
