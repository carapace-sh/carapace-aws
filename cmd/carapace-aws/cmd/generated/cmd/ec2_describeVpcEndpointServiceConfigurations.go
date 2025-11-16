package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcEndpointServiceConfigurationsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoint-service-configurations",
	Short: "Describes the VPC endpoint service configurations in your account (your services).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcEndpointServiceConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcEndpointServiceConfigurationsCmd).Standalone()

		ec2_describeVpcEndpointServiceConfigurationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointServiceConfigurationsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVpcEndpointServiceConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
		ec2_describeVpcEndpointServiceConfigurationsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
		ec2_describeVpcEndpointServiceConfigurationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointServiceConfigurationsCmd.Flags().String("service-ids", "", "The IDs of the endpoint services.")
		ec2_describeVpcEndpointServiceConfigurationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcEndpointServiceConfigurationsCmd)
}
