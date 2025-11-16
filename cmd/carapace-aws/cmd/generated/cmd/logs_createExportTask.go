package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_createExportTaskCmd = &cobra.Command{
	Use:   "create-export-task",
	Short: "Creates an export task so that you can efficiently export data from a log group to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_createExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_createExportTaskCmd).Standalone()

		logs_createExportTaskCmd.Flags().String("destination", "", "The name of S3 bucket for the exported log data.")
		logs_createExportTaskCmd.Flags().String("destination-prefix", "", "The prefix used as the start of the key for every object exported.")
		logs_createExportTaskCmd.Flags().String("from", "", "The start time of the range for the request, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`.")
		logs_createExportTaskCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_createExportTaskCmd.Flags().String("log-stream-name-prefix", "", "Export only log streams that match the provided prefix.")
		logs_createExportTaskCmd.Flags().String("task-name", "", "The name of the export task.")
		logs_createExportTaskCmd.Flags().String("to", "", "The end time of the range for the request, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`.")
		logs_createExportTaskCmd.MarkFlagRequired("destination")
		logs_createExportTaskCmd.MarkFlagRequired("from")
		logs_createExportTaskCmd.MarkFlagRequired("log-group-name")
		logs_createExportTaskCmd.MarkFlagRequired("to")
	})
	logsCmd.AddCommand(logs_createExportTaskCmd)
}
