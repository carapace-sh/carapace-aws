package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listDataSourceSyncJobsCmd = &cobra.Command{
	Use:   "list-data-source-sync-jobs",
	Short: "Get information about an Amazon Q Business data source connector synchronization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listDataSourceSyncJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listDataSourceSyncJobsCmd).Standalone()

		qbusiness_listDataSourceSyncJobsCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application connected to the data source.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("data-source-id", "", "The identifier of the data source connector.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("end-time", "", "The end time of the data source connector sync.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("index-id", "", "The identifier of the index used with the Amazon Q Business data source connector.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("max-results", "", "The maximum number of synchronization jobs to return in the response.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("next-token", "", "If the `maxResults` response was incpmplete because there is more data to retriever, Amazon Q Business returns a pagination token in the response.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("start-time", "", "The start time of the data source connector sync.")
		qbusiness_listDataSourceSyncJobsCmd.Flags().String("status-filter", "", "Only returns synchronization jobs with the `Status` field equal to the specified status.")
		qbusiness_listDataSourceSyncJobsCmd.MarkFlagRequired("application-id")
		qbusiness_listDataSourceSyncJobsCmd.MarkFlagRequired("data-source-id")
		qbusiness_listDataSourceSyncJobsCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listDataSourceSyncJobsCmd)
}
