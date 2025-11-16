package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteFolderMembershipCmd = &cobra.Command{
	Use:   "delete-folder-membership",
	Short: "Removes an asset, such as a dashboard, analysis, or dataset, from a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteFolderMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteFolderMembershipCmd).Standalone()

		quicksight_deleteFolderMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
		quicksight_deleteFolderMembershipCmd.Flags().String("folder-id", "", "The Folder ID.")
		quicksight_deleteFolderMembershipCmd.Flags().String("member-id", "", "The ID of the asset that you want to delete.")
		quicksight_deleteFolderMembershipCmd.Flags().String("member-type", "", "The member type of the asset that you want to delete from a folder.")
		quicksight_deleteFolderMembershipCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteFolderMembershipCmd.MarkFlagRequired("folder-id")
		quicksight_deleteFolderMembershipCmd.MarkFlagRequired("member-id")
		quicksight_deleteFolderMembershipCmd.MarkFlagRequired("member-type")
	})
	quicksightCmd.AddCommand(quicksight_deleteFolderMembershipCmd)
}
