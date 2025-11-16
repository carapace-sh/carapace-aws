package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcEndpointsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoints",
	Short: "Describes your VPC endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcEndpointsCmd).Standalone()

		ec2_describeVpcEndpointsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVpcEndpointsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeVpcEndpointsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ec2_describeVpcEndpointsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointsCmd.Flags().String("vpc-endpoint-ids", "", "The IDs of the VPC endpoints.")
		ec2_describeVpcEndpointsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcEndpointsCmd)
}
