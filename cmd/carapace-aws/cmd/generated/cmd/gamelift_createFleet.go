package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createFleetCmd = &cobra.Command{
	Use:   "create-fleet",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCreates a fleet of compute resources to host your game servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createFleetCmd).Standalone()

		gamelift_createFleetCmd.Flags().String("anywhere-configuration", "", "Amazon GameLift Servers Anywhere configuration options.")
		gamelift_createFleetCmd.Flags().String("build-id", "", "The unique identifier for a custom game server build to be deployed to a fleet with compute type `EC2`.")
		gamelift_createFleetCmd.Flags().String("certificate-configuration", "", "Prompts Amazon GameLift Servers to generate a TLS/SSL certificate for the fleet.")
		gamelift_createFleetCmd.Flags().String("compute-type", "", "The type of compute resource used to host your game servers.")
		gamelift_createFleetCmd.Flags().String("description", "", "A description for the fleet.")
		gamelift_createFleetCmd.Flags().String("ec2-inbound-permissions", "", "The IP address ranges and port settings that allow inbound traffic to access game server processes and other processes on this fleet.")
		gamelift_createFleetCmd.Flags().String("ec2-instance-type", "", "The Amazon GameLift Servers-supported Amazon EC2 instance type to use with managed EC2 fleets.")
		gamelift_createFleetCmd.Flags().String("fleet-type", "", "Indicates whether to use On-Demand or Spot instances for this fleet.")
		gamelift_createFleetCmd.Flags().String("instance-role-arn", "", "A unique identifier for an IAM role that manages access to your Amazon Web Services services.")
		gamelift_createFleetCmd.Flags().String("instance-role-credentials-provider", "", "Prompts Amazon GameLift Servers to generate a shared credentials file for the IAM role that's defined in `InstanceRoleArn`.")
		gamelift_createFleetCmd.Flags().String("locations", "", "A set of remote locations to deploy additional instances to and manage as a multi-location fleet.")
		gamelift_createFleetCmd.Flags().String("log-paths", "", "**This parameter is no longer used.** To specify where Amazon GameLift Servers should store log files once a server process shuts down, use the Amazon GameLift Servers server API `ProcessReady()` and specify one or more directory paths in `logParameters`.")
		gamelift_createFleetCmd.Flags().String("metric-groups", "", "The name of an Amazon Web Services CloudWatch metric group to add this fleet to.")
		gamelift_createFleetCmd.Flags().String("name", "", "A descriptive label that is associated with a fleet.")
		gamelift_createFleetCmd.Flags().String("new-game-session-protection-policy", "", "The status of termination protection for active game sessions on the fleet.")
		gamelift_createFleetCmd.Flags().String("peer-vpc-aws-account-id", "", "Used when peering your Amazon GameLift Servers fleet with a VPC, the unique identifier for the Amazon Web Services account that owns the VPC.")
		gamelift_createFleetCmd.Flags().String("peer-vpc-id", "", "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift Servers fleet.")
		gamelift_createFleetCmd.Flags().String("resource-creation-limit-policy", "", "A policy that limits the number of game sessions that an individual player can create on instances in this fleet within a specified span of time.")
		gamelift_createFleetCmd.Flags().String("runtime-configuration", "", "Instructions for how to launch and run server processes on the fleet.")
		gamelift_createFleetCmd.Flags().String("script-id", "", "The unique identifier for a Realtime configuration script to be deployed to a fleet with compute type `EC2`.")
		gamelift_createFleetCmd.Flags().String("server-launch-parameters", "", "**This parameter is no longer used.** Specify server launch parameters using the `RuntimeConfiguration` parameter.")
		gamelift_createFleetCmd.Flags().String("server-launch-path", "", "**This parameter is no longer used.** Specify a server launch path using the `RuntimeConfiguration` parameter.")
		gamelift_createFleetCmd.Flags().String("tags", "", "A list of labels to assign to the new fleet resource.")
		gamelift_createFleetCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_createFleetCmd)
}
