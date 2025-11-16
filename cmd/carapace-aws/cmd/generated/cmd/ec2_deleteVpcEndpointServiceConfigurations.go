package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpcEndpointServiceConfigurationsCmd = &cobra.Command{
	Use:   "delete-vpc-endpoint-service-configurations",
	Short: "Deletes the specified VPC endpoint service configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpcEndpointServiceConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVpcEndpointServiceConfigurationsCmd).Standalone()

		ec2_deleteVpcEndpointServiceConfigurationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcEndpointServiceConfigurationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcEndpointServiceConfigurationsCmd.Flags().String("service-ids", "", "The IDs of the services.")
		ec2_deleteVpcEndpointServiceConfigurationsCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteVpcEndpointServiceConfigurationsCmd.MarkFlagRequired("service-ids")
	})
	ec2Cmd.AddCommand(ec2_deleteVpcEndpointServiceConfigurationsCmd)
}
