package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listAssetBundleImportJobsCmd = &cobra.Command{
	Use:   "list-asset-bundle-import-jobs",
	Short: "Lists all asset bundle import jobs that have taken place in the last 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listAssetBundleImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listAssetBundleImportJobsCmd).Standalone()

		quicksight_listAssetBundleImportJobsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the import jobs were executed in.")
		quicksight_listAssetBundleImportJobsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listAssetBundleImportJobsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listAssetBundleImportJobsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listAssetBundleImportJobsCmd)
}
