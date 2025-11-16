package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_putClusterCapacityProvidersCmd = &cobra.Command{
	Use:   "put-cluster-capacity-providers",
	Short: "Modifies the available capacity providers and the default capacity provider strategy for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_putClusterCapacityProvidersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_putClusterCapacityProvidersCmd).Standalone()

		ecs_putClusterCapacityProvidersCmd.Flags().String("capacity-providers", "", "The name of one or more capacity providers to associate with the cluster.")
		ecs_putClusterCapacityProvidersCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to modify the capacity provider settings for.")
		ecs_putClusterCapacityProvidersCmd.Flags().String("default-capacity-provider-strategy", "", "The capacity provider strategy to use by default for the cluster.")
		ecs_putClusterCapacityProvidersCmd.MarkFlagRequired("capacity-providers")
		ecs_putClusterCapacityProvidersCmd.MarkFlagRequired("cluster")
		ecs_putClusterCapacityProvidersCmd.MarkFlagRequired("default-capacity-provider-strategy")
	})
	ecsCmd.AddCommand(ecs_putClusterCapacityProvidersCmd)
}
