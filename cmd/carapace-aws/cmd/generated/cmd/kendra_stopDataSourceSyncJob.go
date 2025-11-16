package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_stopDataSourceSyncJobCmd = &cobra.Command{
	Use:   "stop-data-source-sync-job",
	Short: "Stops a synchronization job that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_stopDataSourceSyncJobCmd).Standalone()

	kendra_stopDataSourceSyncJobCmd.Flags().String("id", "", "The identifier of the data source connector for which to stop the synchronization jobs.")
	kendra_stopDataSourceSyncJobCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
	kendra_stopDataSourceSyncJobCmd.MarkFlagRequired("id")
	kendra_stopDataSourceSyncJobCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_stopDataSourceSyncJobCmd)
}
