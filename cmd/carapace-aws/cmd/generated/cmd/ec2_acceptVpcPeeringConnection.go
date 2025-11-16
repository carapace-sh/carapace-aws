package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "accept-vpc-peering-connection",
	Short: "Accept a VPC peering connection request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptVpcPeeringConnectionCmd).Standalone()

	ec2_acceptVpcPeeringConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptVpcPeeringConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptVpcPeeringConnectionCmd.Flags().String("vpc-peering-connection-id", "", "The ID of the VPC peering connection.")
	ec2_acceptVpcPeeringConnectionCmd.Flag("no-dry-run").Hidden = true
	ec2_acceptVpcPeeringConnectionCmd.MarkFlagRequired("vpc-peering-connection-id")
	ec2Cmd.AddCommand(ec2_acceptVpcPeeringConnectionCmd)
}
