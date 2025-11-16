package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listAssetBundleExportJobsCmd = &cobra.Command{
	Use:   "list-asset-bundle-export-jobs",
	Short: "Lists all asset bundle export jobs that have been taken place in the last 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listAssetBundleExportJobsCmd).Standalone()

	quicksight_listAssetBundleExportJobsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the export jobs were executed in.")
	quicksight_listAssetBundleExportJobsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listAssetBundleExportJobsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listAssetBundleExportJobsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listAssetBundleExportJobsCmd)
}
