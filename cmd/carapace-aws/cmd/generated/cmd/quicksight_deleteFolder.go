package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder",
	Short: "Deletes an empty folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteFolderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteFolderCmd).Standalone()

		quicksight_deleteFolderCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
		quicksight_deleteFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
		quicksight_deleteFolderCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteFolderCmd.MarkFlagRequired("folder-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteFolderCmd)
}
