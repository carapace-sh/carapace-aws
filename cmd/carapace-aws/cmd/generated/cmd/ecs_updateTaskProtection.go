package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateTaskProtectionCmd = &cobra.Command{
	Use:   "update-task-protection",
	Short: "Updates the protection status of a task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateTaskProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_updateTaskProtectionCmd).Standalone()

		ecs_updateTaskProtectionCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task sets exist in.")
		ecs_updateTaskProtectionCmd.Flags().String("expires-in-minutes", "", "If you set `protectionEnabled` to `true`, you can specify the duration for task protection in minutes.")
		ecs_updateTaskProtectionCmd.Flags().Bool("no-protection-enabled", false, "Specify `true` to mark a task for protection and `false` to unset protection, making it eligible for termination.")
		ecs_updateTaskProtectionCmd.Flags().Bool("protection-enabled", false, "Specify `true` to mark a task for protection and `false` to unset protection, making it eligible for termination.")
		ecs_updateTaskProtectionCmd.Flags().String("tasks", "", "A list of up to 10 task IDs or full ARN entries.")
		ecs_updateTaskProtectionCmd.MarkFlagRequired("cluster")
		ecs_updateTaskProtectionCmd.Flag("no-protection-enabled").Hidden = true
		ecs_updateTaskProtectionCmd.MarkFlagRequired("no-protection-enabled")
		ecs_updateTaskProtectionCmd.MarkFlagRequired("protection-enabled")
		ecs_updateTaskProtectionCmd.MarkFlagRequired("tasks")
	})
	ecsCmd.AddCommand(ecs_updateTaskProtectionCmd)
}
