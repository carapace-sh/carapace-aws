package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createContainerGroupDefinitionCmd = &cobra.Command{
	Use:   "create-container-group-definition",
	Short: "**This API works with the following fleet types:** Container\n\nCreates a `ContainerGroupDefinition` that describes a set of containers for hosting your game server with Amazon GameLift Servers managed containers hosting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createContainerGroupDefinitionCmd).Standalone()

	gamelift_createContainerGroupDefinitionCmd.Flags().String("container-group-type", "", "The type of container group being defined.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("game-server-container-definition", "", "The definition for the game server container in this group.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("name", "", "A descriptive identifier for the container group definition.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("operating-system", "", "The platform that all containers in the group use.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("support-container-definitions", "", "One or more definition for support containers in this group.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("tags", "", "A list of labels to assign to the container group definition resource.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("total-memory-limit-mebibytes", "", "The maximum amount of memory (in MiB) to allocate to the container group.")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("total-vcpu-limit", "", "The maximum amount of vCPU units to allocate to the container group (1 vCPU is equal to 1024 CPU units).")
	gamelift_createContainerGroupDefinitionCmd.Flags().String("version-description", "", "A description for the initial version of this container group definition.")
	gamelift_createContainerGroupDefinitionCmd.MarkFlagRequired("name")
	gamelift_createContainerGroupDefinitionCmd.MarkFlagRequired("operating-system")
	gamelift_createContainerGroupDefinitionCmd.MarkFlagRequired("total-memory-limit-mebibytes")
	gamelift_createContainerGroupDefinitionCmd.MarkFlagRequired("total-vcpu-limit")
	gameliftCmd.AddCommand(gamelift_createContainerGroupDefinitionCmd)
}
