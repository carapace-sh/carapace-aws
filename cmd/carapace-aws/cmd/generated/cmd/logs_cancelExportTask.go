package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_cancelExportTaskCmd = &cobra.Command{
	Use:   "cancel-export-task",
	Short: "Cancels the specified export task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_cancelExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_cancelExportTaskCmd).Standalone()

		logs_cancelExportTaskCmd.Flags().String("task-id", "", "The ID of the export task.")
		logs_cancelExportTaskCmd.MarkFlagRequired("task-id")
	})
	logsCmd.AddCommand(logs_cancelExportTaskCmd)
}
