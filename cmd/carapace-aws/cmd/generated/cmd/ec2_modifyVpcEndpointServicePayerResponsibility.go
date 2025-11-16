package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcEndpointServicePayerResponsibilityCmd = &cobra.Command{
	Use:   "modify-vpc-endpoint-service-payer-responsibility",
	Short: "Modifies the payer responsibility for your VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcEndpointServicePayerResponsibilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpcEndpointServicePayerResponsibilityCmd).Standalone()

		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.Flags().String("payer-responsibility", "", "The entity that is responsible for the endpoint costs.")
		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.Flags().String("service-id", "", "The ID of the service.")
		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.MarkFlagRequired("payer-responsibility")
		ec2_modifyVpcEndpointServicePayerResponsibilityCmd.MarkFlagRequired("service-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVpcEndpointServicePayerResponsibilityCmd)
}
