package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_batchAddRoleCmd = &cobra.Command{
	Use:   "batch-add-role",
	Short: "Add a role to multiple users or groups in a private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_batchAddRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_batchAddRoleCmd).Standalone()

		repostspace_batchAddRoleCmd.Flags().String("accessor-ids", "", "The user or group accessor identifiers to add the role to.")
		repostspace_batchAddRoleCmd.Flags().String("role", "", "The role to add to the users or groups.")
		repostspace_batchAddRoleCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
		repostspace_batchAddRoleCmd.MarkFlagRequired("accessor-ids")
		repostspace_batchAddRoleCmd.MarkFlagRequired("role")
		repostspace_batchAddRoleCmd.MarkFlagRequired("space-id")
	})
	repostspaceCmd.AddCommand(repostspace_batchAddRoleCmd)
}
