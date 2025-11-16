package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createVpcPeeringAuthorizationCmd = &cobra.Command{
	Use:   "create-vpc-peering-authorization",
	Short: "**This API works with the following fleet types:** EC2\n\nRequests authorization to create or delete a peer connection between the VPC for your Amazon GameLift Servers fleet and a virtual private cloud (VPC) in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createVpcPeeringAuthorizationCmd).Standalone()

	gamelift_createVpcPeeringAuthorizationCmd.Flags().String("game-lift-aws-account-id", "", "A unique identifier for the Amazon Web Services account that you use to manage your Amazon GameLift Servers fleet.")
	gamelift_createVpcPeeringAuthorizationCmd.Flags().String("peer-vpc-id", "", "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift Servers fleet.")
	gamelift_createVpcPeeringAuthorizationCmd.MarkFlagRequired("game-lift-aws-account-id")
	gamelift_createVpcPeeringAuthorizationCmd.MarkFlagRequired("peer-vpc-id")
	gameliftCmd.AddCommand(gamelift_createVpcPeeringAuthorizationCmd)
}
