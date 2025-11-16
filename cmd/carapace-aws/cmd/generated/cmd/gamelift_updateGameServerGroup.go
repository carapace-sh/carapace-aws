package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateGameServerGroupCmd = &cobra.Command{
	Use:   "update-game-server-group",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nUpdates Amazon GameLift Servers FleetIQ-specific properties for a game server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateGameServerGroupCmd).Standalone()

	gamelift_updateGameServerGroupCmd.Flags().String("balancing-strategy", "", "Indicates how Amazon GameLift Servers FleetIQ balances the use of Spot Instances and On-Demand Instances in the game server group.")
	gamelift_updateGameServerGroupCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group.")
	gamelift_updateGameServerGroupCmd.Flags().String("game-server-protection-policy", "", "A flag that indicates whether instances in the game server group are protected from early termination.")
	gamelift_updateGameServerGroupCmd.Flags().String("instance-definitions", "", "An updated list of Amazon EC2 instance types to use in the Auto Scaling group.")
	gamelift_updateGameServerGroupCmd.Flags().String("role-arn", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) for an IAM role that allows Amazon GameLift Servers to access your Amazon EC2 Auto Scaling groups.")
	gamelift_updateGameServerGroupCmd.MarkFlagRequired("game-server-group-name")
	gameliftCmd.AddCommand(gamelift_updateGameServerGroupCmd)
}
