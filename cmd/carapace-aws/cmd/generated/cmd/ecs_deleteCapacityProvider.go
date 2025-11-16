package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteCapacityProviderCmd = &cobra.Command{
	Use:   "delete-capacity-provider",
	Short: "Deletes the specified capacity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteCapacityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_deleteCapacityProviderCmd).Standalone()

		ecs_deleteCapacityProviderCmd.Flags().String("capacity-provider", "", "The short name or full Amazon Resource Name (ARN) of the capacity provider to delete.")
		ecs_deleteCapacityProviderCmd.Flags().String("cluster", "", "The name of the cluster that contains the capacity provider to delete.")
		ecs_deleteCapacityProviderCmd.MarkFlagRequired("capacity-provider")
	})
	ecsCmd.AddCommand(ecs_deleteCapacityProviderCmd)
}
