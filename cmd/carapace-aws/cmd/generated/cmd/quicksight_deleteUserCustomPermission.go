package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteUserCustomPermissionCmd = &cobra.Command{
	Use:   "delete-user-custom-permission",
	Short: "Deletes a custom permissions profile from a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteUserCustomPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteUserCustomPermissionCmd).Standalone()

		quicksight_deleteUserCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the custom permission configuration that you want to delete.")
		quicksight_deleteUserCustomPermissionCmd.Flags().String("namespace", "", "The namespace that the user belongs to.")
		quicksight_deleteUserCustomPermissionCmd.Flags().String("user-name", "", "The username of the user that you want to remove custom permissions from.")
		quicksight_deleteUserCustomPermissionCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteUserCustomPermissionCmd.MarkFlagRequired("namespace")
		quicksight_deleteUserCustomPermissionCmd.MarkFlagRequired("user-name")
	})
	quicksightCmd.AddCommand(quicksight_deleteUserCustomPermissionCmd)
}
