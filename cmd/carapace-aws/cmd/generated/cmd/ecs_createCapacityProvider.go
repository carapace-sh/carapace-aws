package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_createCapacityProviderCmd = &cobra.Command{
	Use:   "create-capacity-provider",
	Short: "Creates a capacity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_createCapacityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_createCapacityProviderCmd).Standalone()

		ecs_createCapacityProviderCmd.Flags().String("auto-scaling-group-provider", "", "The details of the Auto Scaling group for the capacity provider.")
		ecs_createCapacityProviderCmd.Flags().String("cluster", "", "The name of the cluster to associate with the capacity provider.")
		ecs_createCapacityProviderCmd.Flags().String("managed-instances-provider", "", "The configuration for the Amazon ECS Managed Instances provider.")
		ecs_createCapacityProviderCmd.Flags().String("name", "", "The name of the capacity provider.")
		ecs_createCapacityProviderCmd.Flags().String("tags", "", "The metadata that you apply to the capacity provider to categorize and organize them more conveniently.")
		ecs_createCapacityProviderCmd.MarkFlagRequired("name")
	})
	ecsCmd.AddCommand(ecs_createCapacityProviderCmd)
}
