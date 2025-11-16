package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "create-vpc-peering-connection",
	Short: "**This API works with the following fleet types:** EC2\n\nEstablishes a VPC peering connection between a virtual private cloud (VPC) in an Amazon Web Services account with the VPC for your Amazon GameLift Servers fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createVpcPeeringConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createVpcPeeringConnectionCmd).Standalone()

		gamelift_createVpcPeeringConnectionCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet.")
		gamelift_createVpcPeeringConnectionCmd.Flags().String("peer-vpc-aws-account-id", "", "A unique identifier for the Amazon Web Services account with the VPC that you want to peer your Amazon GameLift Servers fleet with.")
		gamelift_createVpcPeeringConnectionCmd.Flags().String("peer-vpc-id", "", "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift Servers fleet.")
		gamelift_createVpcPeeringConnectionCmd.MarkFlagRequired("fleet-id")
		gamelift_createVpcPeeringConnectionCmd.MarkFlagRequired("peer-vpc-aws-account-id")
		gamelift_createVpcPeeringConnectionCmd.MarkFlagRequired("peer-vpc-id")
	})
	gameliftCmd.AddCommand(gamelift_createVpcPeeringConnectionCmd)
}
