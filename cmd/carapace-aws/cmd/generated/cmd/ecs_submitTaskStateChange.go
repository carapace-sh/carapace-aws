package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_submitTaskStateChangeCmd = &cobra.Command{
	Use:   "submit-task-state-change",
	Short: "This action is only used by the Amazon ECS agent, and it is not intended for use outside of the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_submitTaskStateChangeCmd).Standalone()

	ecs_submitTaskStateChangeCmd.Flags().String("attachments", "", "Any attachments associated with the state change request.")
	ecs_submitTaskStateChangeCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the task.")
	ecs_submitTaskStateChangeCmd.Flags().String("containers", "", "Any containers that's associated with the state change request.")
	ecs_submitTaskStateChangeCmd.Flags().String("execution-stopped-at", "", "The Unix timestamp for the time when the task execution stopped.")
	ecs_submitTaskStateChangeCmd.Flags().String("managed-agents", "", "The details for the managed agent that's associated with the task.")
	ecs_submitTaskStateChangeCmd.Flags().String("pull-started-at", "", "The Unix timestamp for the time when the container image pull started.")
	ecs_submitTaskStateChangeCmd.Flags().String("pull-stopped-at", "", "The Unix timestamp for the time when the container image pull completed.")
	ecs_submitTaskStateChangeCmd.Flags().String("reason", "", "The reason for the state change request.")
	ecs_submitTaskStateChangeCmd.Flags().String("status", "", "The status of the state change request.")
	ecs_submitTaskStateChangeCmd.Flags().String("task", "", "The task ID or full ARN of the task in the state change request.")
	ecsCmd.AddCommand(ecs_submitTaskStateChangeCmd)
}
