package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_resumeBatchLoadTaskCmd = &cobra.Command{
	Use:   "resume-batch-load-task",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_resumeBatchLoadTaskCmd).Standalone()

	timestreamWrite_resumeBatchLoadTaskCmd.Flags().String("task-id", "", "The ID of the batch load task to resume.")
	timestreamWrite_resumeBatchLoadTaskCmd.MarkFlagRequired("task-id")
	timestreamWriteCmd.AddCommand(timestreamWrite_resumeBatchLoadTaskCmd)
}
