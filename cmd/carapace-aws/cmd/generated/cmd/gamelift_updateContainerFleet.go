package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateContainerFleetCmd = &cobra.Command{
	Use:   "update-container-fleet",
	Short: "**This API works with the following fleet types:** Container\n\nUpdates the properties of a managed container fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateContainerFleetCmd).Standalone()

	gamelift_updateContainerFleetCmd.Flags().String("deployment-configuration", "", "Instructions for how to deploy updates to a container fleet, if the fleet update initiates a deployment.")
	gamelift_updateContainerFleetCmd.Flags().String("description", "", "A meaningful description of the container fleet.")
	gamelift_updateContainerFleetCmd.Flags().String("fleet-id", "", "A unique identifier for the container fleet to update.")
	gamelift_updateContainerFleetCmd.Flags().String("game-server-container-group-definition-name", "", "The name or ARN value of a new game server container group definition to deploy on the fleet.")
	gamelift_updateContainerFleetCmd.Flags().String("game-server-container-groups-per-instance", "", "The number of times to replicate the game server container group on each fleet instance.")
	gamelift_updateContainerFleetCmd.Flags().String("game-session-creation-limit-policy", "", "A policy that limits the number of game sessions that each individual player can create on instances in this fleet.")
	gamelift_updateContainerFleetCmd.Flags().String("instance-connection-port-range", "", "A revised set of port numbers to open on each fleet instance.")
	gamelift_updateContainerFleetCmd.Flags().String("instance-inbound-permission-authorizations", "", "A set of ports to add to the container fleet's inbound permissions.")
	gamelift_updateContainerFleetCmd.Flags().String("instance-inbound-permission-revocations", "", "A set of ports to remove from the container fleet's inbound permissions.")
	gamelift_updateContainerFleetCmd.Flags().String("log-configuration", "", "The method for collecting container logs for the fleet.")
	gamelift_updateContainerFleetCmd.Flags().String("metric-groups", "", "The name of an Amazon Web Services CloudWatch metric group to add this fleet to.")
	gamelift_updateContainerFleetCmd.Flags().String("new-game-session-protection-policy", "", "The game session protection policy to apply to all new game sessions that are started in this fleet.")
	gamelift_updateContainerFleetCmd.Flags().String("per-instance-container-group-definition-name", "", "The name or ARN value of a new per-instance container group definition to deploy on the fleet.")
	gamelift_updateContainerFleetCmd.Flags().String("remove-attributes", "", "If set, this update removes a fleet's per-instance container group definition.")
	gamelift_updateContainerFleetCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_updateContainerFleetCmd)
}
