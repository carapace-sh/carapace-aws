package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_submitContainerStateChangeCmd = &cobra.Command{
	Use:   "submit-container-state-change",
	Short: "This action is only used by the Amazon ECS agent, and it is not intended for use outside of the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_submitContainerStateChangeCmd).Standalone()

	ecs_submitContainerStateChangeCmd.Flags().String("cluster", "", "The short name or full ARN of the cluster that hosts the container.")
	ecs_submitContainerStateChangeCmd.Flags().String("container-name", "", "The name of the container.")
	ecs_submitContainerStateChangeCmd.Flags().String("exit-code", "", "The exit code that's returned for the state change request.")
	ecs_submitContainerStateChangeCmd.Flags().String("network-bindings", "", "The network bindings of the container.")
	ecs_submitContainerStateChangeCmd.Flags().String("reason", "", "The reason for the state change request.")
	ecs_submitContainerStateChangeCmd.Flags().String("runtime-id", "", "The ID of the Docker container.")
	ecs_submitContainerStateChangeCmd.Flags().String("status", "", "The status of the state change request.")
	ecs_submitContainerStateChangeCmd.Flags().String("task", "", "The task ID or full Amazon Resource Name (ARN) of the task that hosts the container.")
	ecsCmd.AddCommand(ecs_submitContainerStateChangeCmd)
}
