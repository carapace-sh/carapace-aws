package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deregisterContainerInstanceCmd = &cobra.Command{
	Use:   "deregister-container-instance",
	Short: "Deregisters an Amazon ECS container instance from the specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deregisterContainerInstanceCmd).Standalone()

	ecs_deregisterContainerInstanceCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the container instance to deregister.")
	ecs_deregisterContainerInstanceCmd.Flags().String("container-instance", "", "The container instance ID or full ARN of the container instance to deregister.")
	ecs_deregisterContainerInstanceCmd.Flags().String("force", "", "Forces the container instance to be deregistered.")
	ecs_deregisterContainerInstanceCmd.MarkFlagRequired("container-instance")
	ecsCmd.AddCommand(ecs_deregisterContainerInstanceCmd)
}
