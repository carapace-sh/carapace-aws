package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateUserCustomPermissionCmd = &cobra.Command{
	Use:   "update-user-custom-permission",
	Short: "Updates a custom permissions profile for a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateUserCustomPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateUserCustomPermissionCmd).Standalone()

		quicksight_updateUserCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the custom permission configuration that you want to update.")
		quicksight_updateUserCustomPermissionCmd.Flags().String("custom-permissions-name", "", "The name of the custom permissions that you want to update.")
		quicksight_updateUserCustomPermissionCmd.Flags().String("namespace", "", "The namespace that the user belongs to.")
		quicksight_updateUserCustomPermissionCmd.Flags().String("user-name", "", "The username of the user that you want to update custom permissions for.")
		quicksight_updateUserCustomPermissionCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateUserCustomPermissionCmd.MarkFlagRequired("custom-permissions-name")
		quicksight_updateUserCustomPermissionCmd.MarkFlagRequired("namespace")
		quicksight_updateUserCustomPermissionCmd.MarkFlagRequired("user-name")
	})
	quicksightCmd.AddCommand(quicksight_updateUserCustomPermissionCmd)
}
