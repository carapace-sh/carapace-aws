package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createFolderCmd = &cobra.Command{
	Use:   "create-folder",
	Short: "Creates an empty shared folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createFolderCmd).Standalone()

	quicksight_createFolderCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account where you want to create the folder.")
	quicksight_createFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_createFolderCmd.Flags().String("folder-type", "", "The type of folder.")
	quicksight_createFolderCmd.Flags().String("name", "", "The name of the folder.")
	quicksight_createFolderCmd.Flags().String("parent-folder-arn", "", "The Amazon Resource Name (ARN) for the parent folder.")
	quicksight_createFolderCmd.Flags().String("permissions", "", "A structure that describes the principals and the resource-level permissions of a folder.")
	quicksight_createFolderCmd.Flags().String("sharing-model", "", "An optional parameter that determines the sharing scope of the folder.")
	quicksight_createFolderCmd.Flags().String("tags", "", "Tags for the folder.")
	quicksight_createFolderCmd.MarkFlagRequired("aws-account-id")
	quicksight_createFolderCmd.MarkFlagRequired("folder-id")
	quicksightCmd.AddCommand(quicksight_createFolderCmd)
}
