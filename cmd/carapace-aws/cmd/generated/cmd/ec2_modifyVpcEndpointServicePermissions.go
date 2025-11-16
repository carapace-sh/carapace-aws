package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcEndpointServicePermissionsCmd = &cobra.Command{
	Use:   "modify-vpc-endpoint-service-permissions",
	Short: "Modifies the permissions for your VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcEndpointServicePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpcEndpointServicePermissionsCmd).Standalone()

		ec2_modifyVpcEndpointServicePermissionsCmd.Flags().String("add-allowed-principals", "", "The Amazon Resource Names (ARN) of the principals.")
		ec2_modifyVpcEndpointServicePermissionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointServicePermissionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointServicePermissionsCmd.Flags().String("remove-allowed-principals", "", "The Amazon Resource Names (ARN) of the principals.")
		ec2_modifyVpcEndpointServicePermissionsCmd.Flags().String("service-id", "", "The ID of the service.")
		ec2_modifyVpcEndpointServicePermissionsCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVpcEndpointServicePermissionsCmd.MarkFlagRequired("service-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVpcEndpointServicePermissionsCmd)
}
