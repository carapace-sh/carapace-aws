package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listFoldersForResourceCmd = &cobra.Command{
	Use:   "list-folders-for-resource",
	Short: "List all folders that a resource is a member of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listFoldersForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listFoldersForResourceCmd).Standalone()

		quicksight_listFoldersForResourceCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the resource.")
		quicksight_listFoldersForResourceCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listFoldersForResourceCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listFoldersForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) the resource whose folders you need to list.")
		quicksight_listFoldersForResourceCmd.MarkFlagRequired("aws-account-id")
		quicksight_listFoldersForResourceCmd.MarkFlagRequired("resource-arn")
	})
	quicksightCmd.AddCommand(quicksight_listFoldersForResourceCmd)
}
