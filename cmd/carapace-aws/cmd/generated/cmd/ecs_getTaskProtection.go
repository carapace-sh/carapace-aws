package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_getTaskProtectionCmd = &cobra.Command{
	Use:   "get-task-protection",
	Short: "Retrieves the protection status of tasks in an Amazon ECS service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_getTaskProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_getTaskProtectionCmd).Standalone()

		ecs_getTaskProtectionCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task sets exist in.")
		ecs_getTaskProtectionCmd.Flags().String("tasks", "", "A list of up to 100 task IDs or full ARN entries.")
		ecs_getTaskProtectionCmd.MarkFlagRequired("cluster")
	})
	ecsCmd.AddCommand(ecs_getTaskProtectionCmd)
}
