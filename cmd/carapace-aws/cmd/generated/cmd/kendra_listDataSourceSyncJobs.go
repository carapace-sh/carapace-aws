package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listDataSourceSyncJobsCmd = &cobra.Command{
	Use:   "list-data-source-sync-jobs",
	Short: "Gets statistics about synchronizing a data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listDataSourceSyncJobsCmd).Standalone()

	kendra_listDataSourceSyncJobsCmd.Flags().String("id", "", "The identifier of the data source connector.")
	kendra_listDataSourceSyncJobsCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
	kendra_listDataSourceSyncJobsCmd.Flags().String("max-results", "", "The maximum number of synchronization jobs to return in the response.")
	kendra_listDataSourceSyncJobsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listDataSourceSyncJobsCmd.Flags().String("start-time-filter", "", "When specified, the synchronization jobs returned in the list are limited to jobs between the specified dates.")
	kendra_listDataSourceSyncJobsCmd.Flags().String("status-filter", "", "Only returns synchronization jobs with the `Status` field equal to the specified status.")
	kendra_listDataSourceSyncJobsCmd.MarkFlagRequired("id")
	kendra_listDataSourceSyncJobsCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listDataSourceSyncJobsCmd)
}
