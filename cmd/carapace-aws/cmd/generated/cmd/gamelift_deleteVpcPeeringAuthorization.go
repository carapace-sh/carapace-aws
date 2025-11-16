package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteVpcPeeringAuthorizationCmd = &cobra.Command{
	Use:   "delete-vpc-peering-authorization",
	Short: "**This API works with the following fleet types:** EC2\n\nCancels a pending VPC peering authorization for the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteVpcPeeringAuthorizationCmd).Standalone()

	gamelift_deleteVpcPeeringAuthorizationCmd.Flags().String("game-lift-aws-account-id", "", "A unique identifier for the Amazon Web Services account that you use to manage your Amazon GameLift Servers fleet.")
	gamelift_deleteVpcPeeringAuthorizationCmd.Flags().String("peer-vpc-id", "", "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift Servers fleet.")
	gamelift_deleteVpcPeeringAuthorizationCmd.MarkFlagRequired("game-lift-aws-account-id")
	gamelift_deleteVpcPeeringAuthorizationCmd.MarkFlagRequired("peer-vpc-id")
	gameliftCmd.AddCommand(gamelift_deleteVpcPeeringAuthorizationCmd)
}
