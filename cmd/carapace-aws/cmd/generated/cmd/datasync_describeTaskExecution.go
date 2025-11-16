package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeTaskExecutionCmd = &cobra.Command{
	Use:   "describe-task-execution",
	Short: "Provides information about an execution of your DataSync task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeTaskExecutionCmd).Standalone()

	datasync_describeTaskExecutionCmd.Flags().String("task-execution-arn", "", "Specifies the Amazon Resource Name (ARN) of the task execution that you want information about.")
	datasync_describeTaskExecutionCmd.MarkFlagRequired("task-execution-arn")
	datasyncCmd.AddCommand(datasync_describeTaskExecutionCmd)
}
