package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rejectVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "reject-vpc-peering-connection",
	Short: "Rejects a VPC peering connection request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rejectVpcPeeringConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_rejectVpcPeeringConnectionCmd).Standalone()

		ec2_rejectVpcPeeringConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_rejectVpcPeeringConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_rejectVpcPeeringConnectionCmd.Flags().String("vpc-peering-connection-id", "", "The ID of the VPC peering connection.")
		ec2_rejectVpcPeeringConnectionCmd.Flag("no-dry-run").Hidden = true
		ec2_rejectVpcPeeringConnectionCmd.MarkFlagRequired("vpc-peering-connection-id")
	})
	ec2Cmd.AddCommand(ec2_rejectVpcPeeringConnectionCmd)
}
