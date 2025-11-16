package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "delete-vpc-peering-connection",
	Short: "**This API works with the following fleet types:** EC2\n\nRemoves a VPC peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteVpcPeeringConnectionCmd).Standalone()

	gamelift_deleteVpcPeeringConnectionCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet.")
	gamelift_deleteVpcPeeringConnectionCmd.Flags().String("vpc-peering-connection-id", "", "A unique identifier for a VPC peering connection.")
	gamelift_deleteVpcPeeringConnectionCmd.MarkFlagRequired("fleet-id")
	gamelift_deleteVpcPeeringConnectionCmd.MarkFlagRequired("vpc-peering-connection-id")
	gameliftCmd.AddCommand(gamelift_deleteVpcPeeringConnectionCmd)
}
