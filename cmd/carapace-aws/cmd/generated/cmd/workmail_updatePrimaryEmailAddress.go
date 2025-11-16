package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updatePrimaryEmailAddressCmd = &cobra.Command{
	Use:   "update-primary-email-address",
	Short: "Updates the primary email for a user, group, or resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updatePrimaryEmailAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_updatePrimaryEmailAddressCmd).Standalone()

		workmail_updatePrimaryEmailAddressCmd.Flags().String("email", "", "The value of the email to be updated as primary.")
		workmail_updatePrimaryEmailAddressCmd.Flags().String("entity-id", "", "The user, group, or resource to update.")
		workmail_updatePrimaryEmailAddressCmd.Flags().String("organization-id", "", "The organization that contains the user, group, or resource to update.")
		workmail_updatePrimaryEmailAddressCmd.MarkFlagRequired("email")
		workmail_updatePrimaryEmailAddressCmd.MarkFlagRequired("entity-id")
		workmail_updatePrimaryEmailAddressCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_updatePrimaryEmailAddressCmd)
}
