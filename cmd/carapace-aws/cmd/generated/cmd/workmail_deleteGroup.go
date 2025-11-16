package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group from WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deleteGroupCmd).Standalone()

		workmail_deleteGroupCmd.Flags().String("group-id", "", "The identifier of the group to be deleted.")
		workmail_deleteGroupCmd.Flags().String("organization-id", "", "The organization that contains the group.")
		workmail_deleteGroupCmd.MarkFlagRequired("group-id")
		workmail_deleteGroupCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_deleteGroupCmd)
}
