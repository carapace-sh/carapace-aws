package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateTaskSetCmd = &cobra.Command{
	Use:   "update-task-set",
	Short: "Modifies a task set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateTaskSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_updateTaskSetCmd).Standalone()

		ecs_updateTaskSetCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set is found in.")
		ecs_updateTaskSetCmd.Flags().String("scale", "", "A floating-point percentage of the desired number of tasks to place and keep running in the task set.")
		ecs_updateTaskSetCmd.Flags().String("service", "", "The short name or full Amazon Resource Name (ARN) of the service that the task set is found in.")
		ecs_updateTaskSetCmd.Flags().String("task-set", "", "The short name or full Amazon Resource Name (ARN) of the task set to update.")
		ecs_updateTaskSetCmd.MarkFlagRequired("cluster")
		ecs_updateTaskSetCmd.MarkFlagRequired("scale")
		ecs_updateTaskSetCmd.MarkFlagRequired("service")
		ecs_updateTaskSetCmd.MarkFlagRequired("task-set")
	})
	ecsCmd.AddCommand(ecs_updateTaskSetCmd)
}
