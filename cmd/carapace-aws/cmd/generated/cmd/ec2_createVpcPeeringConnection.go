package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "create-vpc-peering-connection",
	Short: "Requests a VPC peering connection between two VPCs: a requester VPC that you own and an accepter VPC with which to create the connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpcPeeringConnectionCmd).Standalone()

	ec2_createVpcPeeringConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpcPeeringConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpcPeeringConnectionCmd.Flags().String("peer-owner-id", "", "The Amazon Web Services account ID of the owner of the accepter VPC.")
	ec2_createVpcPeeringConnectionCmd.Flags().String("peer-region", "", "The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request.")
	ec2_createVpcPeeringConnectionCmd.Flags().String("peer-vpc-id", "", "The ID of the VPC with which you are creating the VPC peering connection.")
	ec2_createVpcPeeringConnectionCmd.Flags().String("tag-specifications", "", "The tags to assign to the peering connection.")
	ec2_createVpcPeeringConnectionCmd.Flags().String("vpc-id", "", "The ID of the requester VPC.")
	ec2_createVpcPeeringConnectionCmd.Flag("no-dry-run").Hidden = true
	ec2_createVpcPeeringConnectionCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_createVpcPeeringConnectionCmd)
}
