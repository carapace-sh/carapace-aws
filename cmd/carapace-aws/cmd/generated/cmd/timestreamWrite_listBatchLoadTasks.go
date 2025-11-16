package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_listBatchLoadTasksCmd = &cobra.Command{
	Use:   "list-batch-load-tasks",
	Short: "Provides a list of batch load tasks, along with the name, status, when the task is resumable until, and other details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_listBatchLoadTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_listBatchLoadTasksCmd).Standalone()

		timestreamWrite_listBatchLoadTasksCmd.Flags().String("max-results", "", "The total number of items to return in the output.")
		timestreamWrite_listBatchLoadTasksCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		timestreamWrite_listBatchLoadTasksCmd.Flags().String("task-status", "", "Status of the batch load task.")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_listBatchLoadTasksCmd)
}
