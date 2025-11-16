package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteRoleCustomPermissionCmd = &cobra.Command{
	Use:   "delete-role-custom-permission",
	Short: "Removes custom permissions from the role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteRoleCustomPermissionCmd).Standalone()

	quicksight_deleteRoleCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
	quicksight_deleteRoleCustomPermissionCmd.Flags().String("namespace", "", "The namespace that includes the role.")
	quicksight_deleteRoleCustomPermissionCmd.Flags().String("role", "", "The role that you want to remove permissions from.")
	quicksight_deleteRoleCustomPermissionCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteRoleCustomPermissionCmd.MarkFlagRequired("namespace")
	quicksight_deleteRoleCustomPermissionCmd.MarkFlagRequired("role")
	quicksightCmd.AddCommand(quicksight_deleteRoleCustomPermissionCmd)
}
