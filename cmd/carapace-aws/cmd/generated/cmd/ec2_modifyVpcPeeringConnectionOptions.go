package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcPeeringConnectionOptionsCmd = &cobra.Command{
	Use:   "modify-vpc-peering-connection-options",
	Short: "Modifies the VPC peering connection options on one side of a VPC peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcPeeringConnectionOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpcPeeringConnectionOptionsCmd).Standalone()

		ec2_modifyVpcPeeringConnectionOptionsCmd.Flags().String("accepter-peering-connection-options", "", "The VPC peering connection options for the accepter VPC.")
		ec2_modifyVpcPeeringConnectionOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcPeeringConnectionOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcPeeringConnectionOptionsCmd.Flags().String("requester-peering-connection-options", "", "The VPC peering connection options for the requester VPC.")
		ec2_modifyVpcPeeringConnectionOptionsCmd.Flags().String("vpc-peering-connection-id", "", "The ID of the VPC peering connection.")
		ec2_modifyVpcPeeringConnectionOptionsCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVpcPeeringConnectionOptionsCmd.MarkFlagRequired("vpc-peering-connection-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVpcPeeringConnectionOptionsCmd)
}
