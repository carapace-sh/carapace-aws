package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_stopTaskCmd = &cobra.Command{
	Use:   "stop-task",
	Short: "Stops a running task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_stopTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_stopTaskCmd).Standalone()

		ecs_stopTaskCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the task to stop.")
		ecs_stopTaskCmd.Flags().String("reason", "", "An optional message specified when a task is stopped.")
		ecs_stopTaskCmd.Flags().String("task", "", "Thefull Amazon Resource Name (ARN) of the task.")
		ecs_stopTaskCmd.MarkFlagRequired("task")
	})
	ecsCmd.AddCommand(ecs_stopTaskCmd)
}
