package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_batchAddChannelRoleToAccessorsCmd = &cobra.Command{
	Use:   "batch-add-channel-role-to-accessors",
	Short: "Add role to multiple users or groups in a private re:Post channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_batchAddChannelRoleToAccessorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_batchAddChannelRoleToAccessorsCmd).Standalone()

		repostspace_batchAddChannelRoleToAccessorsCmd.Flags().String("accessor-ids", "", "The user or group identifiers to add the role to.")
		repostspace_batchAddChannelRoleToAccessorsCmd.Flags().String("channel-id", "", "The unique ID of the private re:Post channel.")
		repostspace_batchAddChannelRoleToAccessorsCmd.Flags().String("channel-role", "", "The channel role to add to the users or groups.")
		repostspace_batchAddChannelRoleToAccessorsCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
		repostspace_batchAddChannelRoleToAccessorsCmd.MarkFlagRequired("accessor-ids")
		repostspace_batchAddChannelRoleToAccessorsCmd.MarkFlagRequired("channel-id")
		repostspace_batchAddChannelRoleToAccessorsCmd.MarkFlagRequired("channel-role")
		repostspace_batchAddChannelRoleToAccessorsCmd.MarkFlagRequired("space-id")
	})
	repostspaceCmd.AddCommand(repostspace_batchAddChannelRoleToAccessorsCmd)
}
