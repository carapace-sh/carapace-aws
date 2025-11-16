package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listFolderMembersCmd = &cobra.Command{
	Use:   "list-folder-members",
	Short: "List all assets (`DASHBOARD`, `ANALYSIS`, and `DATASET`) in a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listFolderMembersCmd).Standalone()

	quicksight_listFolderMembersCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
	quicksight_listFolderMembersCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_listFolderMembersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listFolderMembersCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listFolderMembersCmd.MarkFlagRequired("aws-account-id")
	quicksight_listFolderMembersCmd.MarkFlagRequired("folder-id")
	quicksightCmd.AddCommand(quicksight_listFolderMembersCmd)
}
