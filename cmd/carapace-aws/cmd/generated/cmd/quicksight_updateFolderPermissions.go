package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateFolderPermissionsCmd = &cobra.Command{
	Use:   "update-folder-permissions",
	Short: "Updates permissions of a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateFolderPermissionsCmd).Standalone()

	quicksight_updateFolderPermissionsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder to update.")
	quicksight_updateFolderPermissionsCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_updateFolderPermissionsCmd.Flags().String("grant-permissions", "", "The permissions that you want to grant on a resource.")
	quicksight_updateFolderPermissionsCmd.Flags().String("revoke-permissions", "", "The permissions that you want to revoke from a resource.")
	quicksight_updateFolderPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateFolderPermissionsCmd.MarkFlagRequired("folder-id")
	quicksightCmd.AddCommand(quicksight_updateFolderPermissionsCmd)
}
