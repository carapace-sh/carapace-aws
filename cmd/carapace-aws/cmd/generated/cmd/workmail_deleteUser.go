package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user from WorkMail and all subsequent systems.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteUserCmd).Standalone()

	workmail_deleteUserCmd.Flags().String("organization-id", "", "The organization that contains the user to be deleted.")
	workmail_deleteUserCmd.Flags().String("user-id", "", "The identifier of the user to be deleted.")
	workmail_deleteUserCmd.MarkFlagRequired("organization-id")
	workmail_deleteUserCmd.MarkFlagRequired("user-id")
	workmailCmd.AddCommand(workmail_deleteUserCmd)
}
