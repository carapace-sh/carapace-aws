package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateFleetAttributesCmd = &cobra.Command{
	Use:   "update-fleet-attributes",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nUpdates a fleet's mutable attributes, such as game session protection and resource creation limits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateFleetAttributesCmd).Standalone()

	gamelift_updateFleetAttributesCmd.Flags().String("anywhere-configuration", "", "Amazon GameLift Servers Anywhere configuration options.")
	gamelift_updateFleetAttributesCmd.Flags().String("description", "", "A human-readable description of a fleet.")
	gamelift_updateFleetAttributesCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to update attribute metadata for.")
	gamelift_updateFleetAttributesCmd.Flags().String("metric-groups", "", "The name of a metric group to add this fleet to.")
	gamelift_updateFleetAttributesCmd.Flags().String("name", "", "A descriptive label that is associated with a fleet.")
	gamelift_updateFleetAttributesCmd.Flags().String("new-game-session-protection-policy", "", "The game session protection policy to apply to all new game sessions created in this fleet.")
	gamelift_updateFleetAttributesCmd.Flags().String("resource-creation-limit-policy", "", "Policy settings that limit the number of game sessions an individual player can create over a span of time.")
	gamelift_updateFleetAttributesCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_updateFleetAttributesCmd)
}
