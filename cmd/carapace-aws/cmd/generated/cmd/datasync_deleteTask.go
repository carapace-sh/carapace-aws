package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_deleteTaskCmd = &cobra.Command{
	Use:   "delete-task",
	Short: "Deletes a transfer task resource from DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_deleteTaskCmd).Standalone()

	datasync_deleteTaskCmd.Flags().String("task-arn", "", "Specifies the Amazon Resource Name (ARN) of the task that you want to delete.")
	datasync_deleteTaskCmd.MarkFlagRequired("task-arn")
	datasyncCmd.AddCommand(datasync_deleteTaskCmd)
}
