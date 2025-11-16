package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates data for the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_updateUserCmd).Standalone()

		workmail_updateUserCmd.Flags().String("city", "", "Updates the user's city.")
		workmail_updateUserCmd.Flags().String("company", "", "Updates the user's company.")
		workmail_updateUserCmd.Flags().String("country", "", "Updates the user's country.")
		workmail_updateUserCmd.Flags().String("department", "", "Updates the user's department.")
		workmail_updateUserCmd.Flags().String("display-name", "", "Updates the display name of the user.")
		workmail_updateUserCmd.Flags().String("first-name", "", "Updates the user's first name.")
		workmail_updateUserCmd.Flags().String("hidden-from-global-address-list", "", "If enabled, the user is hidden from the global address list.")
		workmail_updateUserCmd.Flags().String("identity-provider-user-id", "", "User ID from the IAM Identity Center.")
		workmail_updateUserCmd.Flags().String("initials", "", "Updates the user's initials.")
		workmail_updateUserCmd.Flags().String("job-title", "", "Updates the user's job title.")
		workmail_updateUserCmd.Flags().String("last-name", "", "Updates the user's last name.")
		workmail_updateUserCmd.Flags().String("office", "", "Updates the user's office.")
		workmail_updateUserCmd.Flags().String("organization-id", "", "The identifier for the organization under which the user exists.")
		workmail_updateUserCmd.Flags().String("role", "", "Updates the user role.")
		workmail_updateUserCmd.Flags().String("street", "", "Updates the user's street address.")
		workmail_updateUserCmd.Flags().String("telephone", "", "Updates the user's contact details.")
		workmail_updateUserCmd.Flags().String("user-id", "", "The identifier for the user to be updated.")
		workmail_updateUserCmd.Flags().String("zip-code", "", "Updates the user's zip code.")
		workmail_updateUserCmd.MarkFlagRequired("organization-id")
		workmail_updateUserCmd.MarkFlagRequired("user-id")
	})
	workmailCmd.AddCommand(workmail_updateUserCmd)
}
