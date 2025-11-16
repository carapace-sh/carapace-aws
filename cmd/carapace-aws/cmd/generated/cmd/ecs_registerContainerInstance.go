package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_registerContainerInstanceCmd = &cobra.Command{
	Use:   "register-container-instance",
	Short: "This action is only used by the Amazon ECS agent, and it is not intended for use outside of the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_registerContainerInstanceCmd).Standalone()

	ecs_registerContainerInstanceCmd.Flags().String("attributes", "", "The container instance attributes that this container instance supports.")
	ecs_registerContainerInstanceCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to register your container instance with.")
	ecs_registerContainerInstanceCmd.Flags().String("container-instance-arn", "", "The ARN of the container instance (if it was previously registered).")
	ecs_registerContainerInstanceCmd.Flags().String("instance-identity-document", "", "The instance identity document for the EC2 instance to register.")
	ecs_registerContainerInstanceCmd.Flags().String("instance-identity-document-signature", "", "The instance identity document signature for the EC2 instance to register.")
	ecs_registerContainerInstanceCmd.Flags().String("platform-devices", "", "The devices that are available on the container instance.")
	ecs_registerContainerInstanceCmd.Flags().String("tags", "", "The metadata that you apply to the container instance to help you categorize and organize them.")
	ecs_registerContainerInstanceCmd.Flags().String("total-resources", "", "The resources available on the instance.")
	ecs_registerContainerInstanceCmd.Flags().String("version-info", "", "The version information for the Amazon ECS container agent and Docker daemon that runs on the container instance.")
	ecsCmd.AddCommand(ecs_registerContainerInstanceCmd)
}
