package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createFolderMembershipCmd = &cobra.Command{
	Use:   "create-folder-membership",
	Short: "Adds an asset, such as a dashboard, analysis, or dataset into a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createFolderMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createFolderMembershipCmd).Standalone()

		quicksight_createFolderMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
		quicksight_createFolderMembershipCmd.Flags().String("folder-id", "", "The ID of the folder.")
		quicksight_createFolderMembershipCmd.Flags().String("member-id", "", "The ID of the asset that you want to add to the folder.")
		quicksight_createFolderMembershipCmd.Flags().String("member-type", "", "The member type of the asset that you want to add to a folder.")
		quicksight_createFolderMembershipCmd.MarkFlagRequired("aws-account-id")
		quicksight_createFolderMembershipCmd.MarkFlagRequired("folder-id")
		quicksight_createFolderMembershipCmd.MarkFlagRequired("member-id")
		quicksight_createFolderMembershipCmd.MarkFlagRequired("member-type")
	})
	quicksightCmd.AddCommand(quicksight_createFolderMembershipCmd)
}
