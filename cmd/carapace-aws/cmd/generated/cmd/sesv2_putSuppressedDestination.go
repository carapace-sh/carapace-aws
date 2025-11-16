package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putSuppressedDestinationCmd = &cobra.Command{
	Use:   "put-suppressed-destination",
	Short: "Adds an email address to the suppression list for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putSuppressedDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putSuppressedDestinationCmd).Standalone()

		sesv2_putSuppressedDestinationCmd.Flags().String("email-address", "", "The email address that should be added to the suppression list for your account.")
		sesv2_putSuppressedDestinationCmd.Flags().String("reason", "", "The factors that should cause the email address to be added to the suppression list for your account.")
		sesv2_putSuppressedDestinationCmd.MarkFlagRequired("email-address")
		sesv2_putSuppressedDestinationCmd.MarkFlagRequired("reason")
	})
	sesv2Cmd.AddCommand(sesv2_putSuppressedDestinationCmd)
}
