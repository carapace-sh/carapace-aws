package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchFoldersCmd = &cobra.Command{
	Use:   "search-folders",
	Short: "Searches the subfolders in a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchFoldersCmd).Standalone()

	quicksight_searchFoldersCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
	quicksight_searchFoldersCmd.Flags().String("filters", "", "The filters to apply to the search.")
	quicksight_searchFoldersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_searchFoldersCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_searchFoldersCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchFoldersCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchFoldersCmd)
}
