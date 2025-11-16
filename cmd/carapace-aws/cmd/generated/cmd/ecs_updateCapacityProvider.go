package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateCapacityProviderCmd = &cobra.Command{
	Use:   "update-capacity-provider",
	Short: "Modifies the parameters for a capacity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateCapacityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_updateCapacityProviderCmd).Standalone()

		ecs_updateCapacityProviderCmd.Flags().String("auto-scaling-group-provider", "", "An object that represent the parameters to update for the Auto Scaling group capacity provider.")
		ecs_updateCapacityProviderCmd.Flags().String("cluster", "", "The name of the cluster that contains the capacity provider to update.")
		ecs_updateCapacityProviderCmd.Flags().String("managed-instances-provider", "", "The updated configuration for the Amazon ECS Managed Instances provider.")
		ecs_updateCapacityProviderCmd.Flags().String("name", "", "The name of the capacity provider to update.")
		ecs_updateCapacityProviderCmd.MarkFlagRequired("name")
	})
	ecsCmd.AddCommand(ecs_updateCapacityProviderCmd)
}
