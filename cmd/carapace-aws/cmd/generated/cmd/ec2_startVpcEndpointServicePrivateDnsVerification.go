package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_startVpcEndpointServicePrivateDnsVerificationCmd = &cobra.Command{
	Use:   "start-vpc-endpoint-service-private-dns-verification",
	Short: "Initiates the verification process to prove that the service provider owns the private DNS name domain for the endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_startVpcEndpointServicePrivateDnsVerificationCmd).Standalone()

	ec2_startVpcEndpointServicePrivateDnsVerificationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_startVpcEndpointServicePrivateDnsVerificationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_startVpcEndpointServicePrivateDnsVerificationCmd.Flags().String("service-id", "", "The ID of the endpoint service.")
	ec2_startVpcEndpointServicePrivateDnsVerificationCmd.Flag("no-dry-run").Hidden = true
	ec2_startVpcEndpointServicePrivateDnsVerificationCmd.MarkFlagRequired("service-id")
	ec2Cmd.AddCommand(ec2_startVpcEndpointServicePrivateDnsVerificationCmd)
}
