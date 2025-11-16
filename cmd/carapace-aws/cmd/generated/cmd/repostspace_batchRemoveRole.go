package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_batchRemoveRoleCmd = &cobra.Command{
	Use:   "batch-remove-role",
	Short: "Remove a role from multiple users or groups in a private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_batchRemoveRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_batchRemoveRoleCmd).Standalone()

		repostspace_batchRemoveRoleCmd.Flags().String("accessor-ids", "", "The user or group accessor identifiers to remove the role from.")
		repostspace_batchRemoveRoleCmd.Flags().String("role", "", "The role to remove from the users or groups.")
		repostspace_batchRemoveRoleCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
		repostspace_batchRemoveRoleCmd.MarkFlagRequired("accessor-ids")
		repostspace_batchRemoveRoleCmd.MarkFlagRequired("role")
		repostspace_batchRemoveRoleCmd.MarkFlagRequired("space-id")
	})
	repostspaceCmd.AddCommand(repostspace_batchRemoveRoleCmd)
}
