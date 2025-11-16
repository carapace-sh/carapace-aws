package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createGameServerGroupCmd = &cobra.Command{
	Use:   "create-game-server-group",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nCreates a Amazon GameLift Servers FleetIQ game server group for managing game hosting on a collection of Amazon Elastic Compute Cloud instances for game hosting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createGameServerGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createGameServerGroupCmd).Standalone()

		gamelift_createGameServerGroupCmd.Flags().String("auto-scaling-policy", "", "Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting.")
		gamelift_createGameServerGroupCmd.Flags().String("balancing-strategy", "", "Indicates how Amazon GameLift Servers FleetIQ balances the use of Spot Instances and On-Demand Instances in the game server group.")
		gamelift_createGameServerGroupCmd.Flags().String("game-server-group-name", "", "An identifier for the new game server group.")
		gamelift_createGameServerGroupCmd.Flags().String("game-server-protection-policy", "", "A flag that indicates whether instances in the game server group are protected from early termination.")
		gamelift_createGameServerGroupCmd.Flags().String("instance-definitions", "", "The Amazon EC2 instance types and sizes to use in the Auto Scaling group.")
		gamelift_createGameServerGroupCmd.Flags().String("launch-template", "", "The Amazon EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.")
		gamelift_createGameServerGroupCmd.Flags().String("max-size", "", "The maximum number of instances allowed in the Amazon EC2 Auto Scaling group.")
		gamelift_createGameServerGroupCmd.Flags().String("min-size", "", "The minimum number of instances allowed in the Amazon EC2 Auto Scaling group.")
		gamelift_createGameServerGroupCmd.Flags().String("role-arn", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) for an IAM role that allows Amazon GameLift Servers to access your Amazon EC2 Auto Scaling groups.")
		gamelift_createGameServerGroupCmd.Flags().String("tags", "", "A list of labels to assign to the new game server group resource.")
		gamelift_createGameServerGroupCmd.Flags().String("vpc-subnets", "", "A list of virtual private cloud (VPC) subnets to use with instances in the game server group.")
		gamelift_createGameServerGroupCmd.MarkFlagRequired("game-server-group-name")
		gamelift_createGameServerGroupCmd.MarkFlagRequired("instance-definitions")
		gamelift_createGameServerGroupCmd.MarkFlagRequired("launch-template")
		gamelift_createGameServerGroupCmd.MarkFlagRequired("max-size")
		gamelift_createGameServerGroupCmd.MarkFlagRequired("min-size")
		gamelift_createGameServerGroupCmd.MarkFlagRequired("role-arn")
	})
	gameliftCmd.AddCommand(gamelift_createGameServerGroupCmd)
}
