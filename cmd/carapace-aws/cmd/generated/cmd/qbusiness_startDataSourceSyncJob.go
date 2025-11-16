package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_startDataSourceSyncJobCmd = &cobra.Command{
	Use:   "start-data-source-sync-job",
	Short: "Starts a data source connector synchronization job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_startDataSourceSyncJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_startDataSourceSyncJobCmd).Standalone()

		qbusiness_startDataSourceSyncJobCmd.Flags().String("application-id", "", "The identifier of Amazon Q Business application the data source is connected to.")
		qbusiness_startDataSourceSyncJobCmd.Flags().String("data-source-id", "", "The identifier of the data source connector.")
		qbusiness_startDataSourceSyncJobCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
		qbusiness_startDataSourceSyncJobCmd.MarkFlagRequired("application-id")
		qbusiness_startDataSourceSyncJobCmd.MarkFlagRequired("data-source-id")
		qbusiness_startDataSourceSyncJobCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_startDataSourceSyncJobCmd)
}
