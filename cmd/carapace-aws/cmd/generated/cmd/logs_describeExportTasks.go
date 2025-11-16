package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeExportTasksCmd = &cobra.Command{
	Use:   "describe-export-tasks",
	Short: "Lists the specified export tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeExportTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_describeExportTasksCmd).Standalone()

		logs_describeExportTasksCmd.Flags().String("limit", "", "The maximum number of items returned.")
		logs_describeExportTasksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		logs_describeExportTasksCmd.Flags().String("status-code", "", "The status code of the export task.")
		logs_describeExportTasksCmd.Flags().String("task-id", "", "The ID of the export task.")
	})
	logsCmd.AddCommand(logs_describeExportTasksCmd)
}
