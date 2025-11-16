package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcEndpointServicePermissionsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoint-service-permissions",
	Short: "Describes the principals (service consumers) that are permitted to discover your VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcEndpointServicePermissionsCmd).Standalone()

	ec2_describeVpcEndpointServicePermissionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcEndpointServicePermissionsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeVpcEndpointServicePermissionsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeVpcEndpointServicePermissionsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeVpcEndpointServicePermissionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcEndpointServicePermissionsCmd.Flags().String("service-id", "", "The ID of the service.")
	ec2_describeVpcEndpointServicePermissionsCmd.Flag("no-dry-run").Hidden = true
	ec2_describeVpcEndpointServicePermissionsCmd.MarkFlagRequired("service-id")
	ec2Cmd.AddCommand(ec2_describeVpcEndpointServicePermissionsCmd)
}
