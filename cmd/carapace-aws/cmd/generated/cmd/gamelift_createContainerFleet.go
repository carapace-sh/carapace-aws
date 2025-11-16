package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createContainerFleetCmd = &cobra.Command{
	Use:   "create-container-fleet",
	Short: "**This API works with the following fleet types:** Container\n\nCreates a managed fleet of Amazon Elastic Compute Cloud (Amazon EC2) instances to host your containerized game servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createContainerFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createContainerFleetCmd).Standalone()

		gamelift_createContainerFleetCmd.Flags().String("billing-type", "", "Indicates whether to use On-Demand or Spot instances for this fleet.")
		gamelift_createContainerFleetCmd.Flags().String("description", "", "A meaningful description of the container fleet.")
		gamelift_createContainerFleetCmd.Flags().String("fleet-role-arn", "", "The unique identifier for an Identity and Access Management (IAM) role with permissions to run your containers on resources that are managed by Amazon GameLift Servers.")
		gamelift_createContainerFleetCmd.Flags().String("game-server-container-group-definition-name", "", "A container group definition resource that describes how to deploy containers with your game server build and support software onto each fleet instance.")
		gamelift_createContainerFleetCmd.Flags().String("game-server-container-groups-per-instance", "", "The number of times to replicate the game server container group on each fleet instance.")
		gamelift_createContainerFleetCmd.Flags().String("game-session-creation-limit-policy", "", "A policy that limits the number of game sessions that each individual player can create on instances in this fleet.")
		gamelift_createContainerFleetCmd.Flags().String("instance-connection-port-range", "", "The set of port numbers to open on each fleet instance.")
		gamelift_createContainerFleetCmd.Flags().String("instance-inbound-permissions", "", "The IP address ranges and port settings that allow inbound traffic to access game server processes and other processes on this fleet.")
		gamelift_createContainerFleetCmd.Flags().String("instance-type", "", "The Amazon EC2 instance type to use for all instances in the fleet.")
		gamelift_createContainerFleetCmd.Flags().String("locations", "", "A set of locations to deploy container fleet instances to.")
		gamelift_createContainerFleetCmd.Flags().String("log-configuration", "", "A method for collecting container logs for the fleet.")
		gamelift_createContainerFleetCmd.Flags().String("metric-groups", "", "The name of an Amazon Web Services CloudWatch metric group to add this fleet to.")
		gamelift_createContainerFleetCmd.Flags().String("new-game-session-protection-policy", "", "Determines whether Amazon GameLift Servers can shut down game sessions on the fleet that are actively running and hosting players.")
		gamelift_createContainerFleetCmd.Flags().String("per-instance-container-group-definition-name", "", "The name of a container group definition resource that describes a set of axillary software.")
		gamelift_createContainerFleetCmd.Flags().String("tags", "", "A list of labels to assign to the new fleet resource.")
		gamelift_createContainerFleetCmd.MarkFlagRequired("fleet-role-arn")
	})
	gameliftCmd.AddCommand(gamelift_createContainerFleetCmd)
}
