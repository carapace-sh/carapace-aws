package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_stopDataSourceSyncJobCmd = &cobra.Command{
	Use:   "stop-data-source-sync-job",
	Short: "Stops an Amazon Q Business data source connector synchronization job already in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_stopDataSourceSyncJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_stopDataSourceSyncJobCmd).Standalone()

		qbusiness_stopDataSourceSyncJobCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application that the data source is connected to.")
		qbusiness_stopDataSourceSyncJobCmd.Flags().String("data-source-id", "", "The identifier of the data source connector.")
		qbusiness_stopDataSourceSyncJobCmd.Flags().String("index-id", "", "The identifier of the index used with the Amazon Q Business data source connector.")
		qbusiness_stopDataSourceSyncJobCmd.MarkFlagRequired("application-id")
		qbusiness_stopDataSourceSyncJobCmd.MarkFlagRequired("data-source-id")
		qbusiness_stopDataSourceSyncJobCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_stopDataSourceSyncJobCmd)
}
