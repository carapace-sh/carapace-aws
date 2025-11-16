package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateContainerInstancesStateCmd = &cobra.Command{
	Use:   "update-container-instances-state",
	Short: "Modifies the status of an Amazon ECS container instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateContainerInstancesStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_updateContainerInstancesStateCmd).Standalone()

		ecs_updateContainerInstancesStateCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the container instance to update.")
		ecs_updateContainerInstancesStateCmd.Flags().String("container-instances", "", "A list of up to 10 container instance IDs or full ARN entries.")
		ecs_updateContainerInstancesStateCmd.Flags().String("status", "", "The container instance state to update the container instance with.")
		ecs_updateContainerInstancesStateCmd.MarkFlagRequired("container-instances")
		ecs_updateContainerInstancesStateCmd.MarkFlagRequired("status")
	})
	ecsCmd.AddCommand(ecs_updateContainerInstancesStateCmd)
}
