package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateContainerGroupDefinitionCmd = &cobra.Command{
	Use:   "update-container-group-definition",
	Short: "**This API works with the following fleet types:** Container\n\nUpdates properties in an existing container group definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateContainerGroupDefinitionCmd).Standalone()

	gamelift_updateContainerGroupDefinitionCmd.Flags().String("game-server-container-definition", "", "An updated definition for the game server container in this group.")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("name", "", "A descriptive identifier for the container group definition.")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("operating-system", "", "The platform that all containers in the group use.")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("source-version-number", "", "The container group definition version to update.")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("support-container-definitions", "", "One or more definitions for support containers in this group.")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("total-memory-limit-mebibytes", "", "The maximum amount of memory (in MiB) to allocate to the container group.")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("total-vcpu-limit", "", "The maximum amount of vCPU units to allocate to the container group (1 vCPU is equal to 1024 CPU units).")
	gamelift_updateContainerGroupDefinitionCmd.Flags().String("version-description", "", "A description for this update to the container group definition.")
	gamelift_updateContainerGroupDefinitionCmd.MarkFlagRequired("name")
	gameliftCmd.AddCommand(gamelift_updateContainerGroupDefinitionCmd)
}
