package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_describeBatchLoadTaskCmd = &cobra.Command{
	Use:   "describe-batch-load-task",
	Short: "Returns information about the batch load task, including configurations, mappings, progress, and other details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_describeBatchLoadTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_describeBatchLoadTaskCmd).Standalone()

		timestreamWrite_describeBatchLoadTaskCmd.Flags().String("task-id", "", "The ID of the batch load task.")
		timestreamWrite_describeBatchLoadTaskCmd.MarkFlagRequired("task-id")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_describeBatchLoadTaskCmd)
}
