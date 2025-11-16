package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateRoleCustomPermissionCmd = &cobra.Command{
	Use:   "update-role-custom-permission",
	Short: "Updates the custom permissions that are associated with a role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateRoleCustomPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateRoleCustomPermissionCmd).Standalone()

		quicksight_updateRoleCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to create a group in.")
		quicksight_updateRoleCustomPermissionCmd.Flags().String("custom-permissions-name", "", "The name of the custom permission that you want to update the role with.")
		quicksight_updateRoleCustomPermissionCmd.Flags().String("namespace", "", "The namespace that contains the role that you want to update.")
		quicksight_updateRoleCustomPermissionCmd.Flags().String("role", "", "The name of role tht you want to update.")
		quicksight_updateRoleCustomPermissionCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateRoleCustomPermissionCmd.MarkFlagRequired("custom-permissions-name")
		quicksight_updateRoleCustomPermissionCmd.MarkFlagRequired("namespace")
		quicksight_updateRoleCustomPermissionCmd.MarkFlagRequired("role")
	})
	quicksightCmd.AddCommand(quicksight_updateRoleCustomPermissionCmd)
}
