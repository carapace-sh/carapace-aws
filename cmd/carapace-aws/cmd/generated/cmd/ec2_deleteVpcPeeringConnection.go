package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "delete-vpc-peering-connection",
	Short: "Deletes a VPC peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpcPeeringConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVpcPeeringConnectionCmd).Standalone()

		ec2_deleteVpcPeeringConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcPeeringConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcPeeringConnectionCmd.Flags().String("vpc-peering-connection-id", "", "The ID of the VPC peering connection.")
		ec2_deleteVpcPeeringConnectionCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteVpcPeeringConnectionCmd.MarkFlagRequired("vpc-peering-connection-id")
	})
	ec2Cmd.AddCommand(ec2_deleteVpcPeeringConnectionCmd)
}
