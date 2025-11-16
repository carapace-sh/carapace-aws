package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateFolderCmd = &cobra.Command{
	Use:   "update-folder",
	Short: "Updates the name of a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateFolderCmd).Standalone()

	quicksight_updateFolderCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder to update.")
	quicksight_updateFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_updateFolderCmd.Flags().String("name", "", "The name of the folder.")
	quicksight_updateFolderCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateFolderCmd.MarkFlagRequired("folder-id")
	quicksight_updateFolderCmd.MarkFlagRequired("name")
	quicksightCmd.AddCommand(quicksight_updateFolderCmd)
}
