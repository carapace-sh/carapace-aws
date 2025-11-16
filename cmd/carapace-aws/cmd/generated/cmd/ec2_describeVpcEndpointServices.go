package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcEndpointServicesCmd = &cobra.Command{
	Use:   "describe-vpc-endpoint-services",
	Short: "Describes available services to which you can create a VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcEndpointServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcEndpointServicesCmd).Standalone()

		ec2_describeVpcEndpointServicesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointServicesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeVpcEndpointServicesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeVpcEndpointServicesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ec2_describeVpcEndpointServicesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcEndpointServicesCmd.Flags().String("service-names", "", "The service names.")
		ec2_describeVpcEndpointServicesCmd.Flags().String("service-regions", "", "The service Regions.")
		ec2_describeVpcEndpointServicesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcEndpointServicesCmd)
}
