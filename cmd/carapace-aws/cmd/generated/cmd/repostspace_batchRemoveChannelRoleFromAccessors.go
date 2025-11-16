package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_batchRemoveChannelRoleFromAccessorsCmd = &cobra.Command{
	Use:   "batch-remove-channel-role-from-accessors",
	Short: "Remove a role from multiple users or groups in a private re:Post channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_batchRemoveChannelRoleFromAccessorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_batchRemoveChannelRoleFromAccessorsCmd).Standalone()

		repostspace_batchRemoveChannelRoleFromAccessorsCmd.Flags().String("accessor-ids", "", "The users or groups identifiers to remove the role from.")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.Flags().String("channel-id", "", "The unique ID of the private re:Post channel.")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.Flags().String("channel-role", "", "The channel role to remove from the users or groups.")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.MarkFlagRequired("accessor-ids")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.MarkFlagRequired("channel-id")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.MarkFlagRequired("channel-role")
		repostspace_batchRemoveChannelRoleFromAccessorsCmd.MarkFlagRequired("space-id")
	})
	repostspaceCmd.AddCommand(repostspace_batchRemoveChannelRoleFromAccessorsCmd)
}
