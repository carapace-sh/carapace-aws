package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateContainerAgentCmd = &cobra.Command{
	Use:   "update-container-agent",
	Short: "Updates the Amazon ECS container agent on a specified container instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateContainerAgentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_updateContainerAgentCmd).Standalone()

		ecs_updateContainerAgentCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that your container instance is running on.")
		ecs_updateContainerAgentCmd.Flags().String("container-instance", "", "The container instance ID or full ARN entries for the container instance where you would like to update the Amazon ECS container agent.")
		ecs_updateContainerAgentCmd.MarkFlagRequired("container-instance")
	})
	ecsCmd.AddCommand(ecs_updateContainerAgentCmd)
}
