package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_startDataSourceSyncJobCmd = &cobra.Command{
	Use:   "start-data-source-sync-job",
	Short: "Starts a synchronization job for a data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_startDataSourceSyncJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_startDataSourceSyncJobCmd).Standalone()

		kendra_startDataSourceSyncJobCmd.Flags().String("id", "", "The identifier of the data source connector to synchronize.")
		kendra_startDataSourceSyncJobCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
		kendra_startDataSourceSyncJobCmd.MarkFlagRequired("id")
		kendra_startDataSourceSyncJobCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_startDataSourceSyncJobCmd)
}
