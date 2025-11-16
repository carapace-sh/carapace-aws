package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_startImportTaskCmd = &cobra.Command{
	Use:   "start-import-task",
	Short: "Import data into existing Neptune Analytics graph from Amazon Simple Storage Service (S3).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_startImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_startImportTaskCmd).Standalone()

		neptuneGraph_startImportTaskCmd.Flags().String("blank-node-handling", "", "The method to handle blank nodes in the dataset.")
		neptuneGraph_startImportTaskCmd.Flags().Bool("fail-on-error", false, "If set to true, the task halts when an import error is encountered.")
		neptuneGraph_startImportTaskCmd.Flags().String("format", "", "Specifies the format of Amazon S3 data to be imported.")
		neptuneGraph_startImportTaskCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_startImportTaskCmd.Flags().String("import-options", "", "")
		neptuneGraph_startImportTaskCmd.Flags().Bool("no-fail-on-error", false, "If set to true, the task halts when an import error is encountered.")
		neptuneGraph_startImportTaskCmd.Flags().String("parquet-type", "", "The parquet type of the import task.")
		neptuneGraph_startImportTaskCmd.Flags().String("role-arn", "", "The ARN of the IAM role that will allow access to the data that is to be imported.")
		neptuneGraph_startImportTaskCmd.Flags().String("source", "", "A URL identifying the location of the data to be imported.")
		neptuneGraph_startImportTaskCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_startImportTaskCmd.Flag("no-fail-on-error").Hidden = true
		neptuneGraph_startImportTaskCmd.MarkFlagRequired("role-arn")
		neptuneGraph_startImportTaskCmd.MarkFlagRequired("source")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_startImportTaskCmd)
}
