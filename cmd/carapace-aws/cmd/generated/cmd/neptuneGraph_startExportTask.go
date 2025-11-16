package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_startExportTaskCmd = &cobra.Command{
	Use:   "start-export-task",
	Short: "Export data from an existing Neptune Analytics graph to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_startExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_startExportTaskCmd).Standalone()

		neptuneGraph_startExportTaskCmd.Flags().String("destination", "", "The Amazon S3 URI where data will be exported to.")
		neptuneGraph_startExportTaskCmd.Flags().String("export-filter", "", "The export filter of the export task.")
		neptuneGraph_startExportTaskCmd.Flags().String("format", "", "The format of the export task.")
		neptuneGraph_startExportTaskCmd.Flags().String("graph-identifier", "", "The source graph identifier of the export task.")
		neptuneGraph_startExportTaskCmd.Flags().String("kms-key-identifier", "", "The KMS key identifier of the export task.")
		neptuneGraph_startExportTaskCmd.Flags().String("parquet-type", "", "The parquet type of the export task.")
		neptuneGraph_startExportTaskCmd.Flags().String("role-arn", "", "The ARN of the IAM role that will allow data to be exported to the destination.")
		neptuneGraph_startExportTaskCmd.Flags().String("tags", "", "Tags to be applied to the export task.")
		neptuneGraph_startExportTaskCmd.MarkFlagRequired("destination")
		neptuneGraph_startExportTaskCmd.MarkFlagRequired("format")
		neptuneGraph_startExportTaskCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_startExportTaskCmd.MarkFlagRequired("kms-key-identifier")
		neptuneGraph_startExportTaskCmd.MarkFlagRequired("role-arn")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_startExportTaskCmd)
}
