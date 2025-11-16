package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcEndpointAssociationsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoint-associations",
	Short: "Describes the VPC resources, VPC endpoint services, Amazon Lattice services, or service networks associated with the VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcEndpointAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcEndpointAssociationsCmd).Standalone()

		ec2_describeVpcEndpointAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointAssociationsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVpcEndpointAssociationsCmd.Flags().String("max-results", "", "The maximum page size.")
		ec2_describeVpcEndpointAssociationsCmd.Flags().String("next-token", "", "The pagination token.")
		ec2_describeVpcEndpointAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointAssociationsCmd.Flags().String("vpc-endpoint-ids", "", "The IDs of the VPC endpoints.")
		ec2_describeVpcEndpointAssociationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcEndpointAssociationsCmd)
}
