package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_createBatchLoadTaskCmd = &cobra.Command{
	Use:   "create-batch-load-task",
	Short: "Creates a new Timestream batch load task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_createBatchLoadTaskCmd).Standalone()

	timestreamWrite_createBatchLoadTaskCmd.Flags().String("client-token", "", "")
	timestreamWrite_createBatchLoadTaskCmd.Flags().String("data-model-configuration", "", "")
	timestreamWrite_createBatchLoadTaskCmd.Flags().String("data-source-configuration", "", "Defines configuration details about the data source for a batch load task.")
	timestreamWrite_createBatchLoadTaskCmd.Flags().String("record-version", "", "")
	timestreamWrite_createBatchLoadTaskCmd.Flags().String("report-configuration", "", "")
	timestreamWrite_createBatchLoadTaskCmd.Flags().String("target-database-name", "", "Target Timestream database for a batch load task.")
	timestreamWrite_createBatchLoadTaskCmd.Flags().String("target-table-name", "", "Target Timestream table for a batch load task.")
	timestreamWrite_createBatchLoadTaskCmd.MarkFlagRequired("data-source-configuration")
	timestreamWrite_createBatchLoadTaskCmd.MarkFlagRequired("report-configuration")
	timestreamWrite_createBatchLoadTaskCmd.MarkFlagRequired("target-database-name")
	timestreamWrite_createBatchLoadTaskCmd.MarkFlagRequired("target-table-name")
	timestreamWriteCmd.AddCommand(timestreamWrite_createBatchLoadTaskCmd)
}
