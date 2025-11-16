package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteTaskSetCmd = &cobra.Command{
	Use:   "delete-task-set",
	Short: "Deletes a specified task set within a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteTaskSetCmd).Standalone()

	ecs_deleteTaskSetCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set found in to delete.")
	ecs_deleteTaskSetCmd.Flags().String("force", "", "If `true`, you can delete a task set even if it hasn't been scaled down to zero.")
	ecs_deleteTaskSetCmd.Flags().String("service", "", "The short name or full Amazon Resource Name (ARN) of the service that hosts the task set to delete.")
	ecs_deleteTaskSetCmd.Flags().String("task-set", "", "The task set ID or full Amazon Resource Name (ARN) of the task set to delete.")
	ecs_deleteTaskSetCmd.MarkFlagRequired("cluster")
	ecs_deleteTaskSetCmd.MarkFlagRequired("service")
	ecs_deleteTaskSetCmd.MarkFlagRequired("task-set")
	ecsCmd.AddCommand(ecs_deleteTaskSetCmd)
}
