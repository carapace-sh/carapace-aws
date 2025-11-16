package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_discoverPollEndpointCmd = &cobra.Command{
	Use:   "discover-poll-endpoint",
	Short: "This action is only used by the Amazon ECS agent, and it is not intended for use outside of the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_discoverPollEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_discoverPollEndpointCmd).Standalone()

		ecs_discoverPollEndpointCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that the container instance belongs to.")
		ecs_discoverPollEndpointCmd.Flags().String("container-instance", "", "The container instance ID or full ARN of the container instance.")
	})
	ecsCmd.AddCommand(ecs_discoverPollEndpointCmd)
}
