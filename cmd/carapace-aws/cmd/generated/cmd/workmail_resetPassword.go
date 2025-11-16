package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_resetPasswordCmd = &cobra.Command{
	Use:   "reset-password",
	Short: "Allows the administrator to reset the password for a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_resetPasswordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_resetPasswordCmd).Standalone()

		workmail_resetPasswordCmd.Flags().String("organization-id", "", "The identifier of the organization that contains the user for which the password is reset.")
		workmail_resetPasswordCmd.Flags().String("password", "", "The new password for the user.")
		workmail_resetPasswordCmd.Flags().String("user-id", "", "The identifier of the user for whom the password is reset.")
		workmail_resetPasswordCmd.MarkFlagRequired("organization-id")
		workmail_resetPasswordCmd.MarkFlagRequired("password")
		workmail_resetPasswordCmd.MarkFlagRequired("user-id")
	})
	workmailCmd.AddCommand(workmail_resetPasswordCmd)
}
