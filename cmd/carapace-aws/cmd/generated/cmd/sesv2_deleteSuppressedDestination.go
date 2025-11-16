package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteSuppressedDestinationCmd = &cobra.Command{
	Use:   "delete-suppressed-destination",
	Short: "Removes an email address from the suppression list for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteSuppressedDestinationCmd).Standalone()

	sesv2_deleteSuppressedDestinationCmd.Flags().String("email-address", "", "The suppressed email destination to remove from the account suppression list.")
	sesv2_deleteSuppressedDestinationCmd.MarkFlagRequired("email-address")
	sesv2Cmd.AddCommand(sesv2_deleteSuppressedDestinationCmd)
}
