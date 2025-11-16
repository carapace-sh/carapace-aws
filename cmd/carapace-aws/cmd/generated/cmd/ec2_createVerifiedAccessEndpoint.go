package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVerifiedAccessEndpointCmd = &cobra.Command{
	Use:   "create-verified-access-endpoint",
	Short: "An Amazon Web Services Verified Access endpoint is where you define your application along with an optional endpoint-level access policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVerifiedAccessEndpointCmd).Standalone()

	ec2_createVerifiedAccessEndpointCmd.Flags().String("application-domain", "", "The DNS name for users to reach your application.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("attachment-type", "", "The type of attachment.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("cidr-options", "", "The CIDR options.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("description", "", "A description for the Verified Access endpoint.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("domain-certificate-arn", "", "The ARN of the public TLS/SSL certificate in Amazon Web Services Certificate Manager to associate with the endpoint.")
	ec2_createVerifiedAccessEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("endpoint-domain-prefix", "", "A custom identifier that is prepended to the DNS name that is generated for the endpoint.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("endpoint-type", "", "The type of Verified Access endpoint to create.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("load-balancer-options", "", "The load balancer details.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("network-interface-options", "", "The network interface details.")
	ec2_createVerifiedAccessEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("policy-document", "", "The Verified Access policy document.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("rds-options", "", "The RDS details.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("security-group-ids", "", "The IDs of the security groups to associate with the Verified Access endpoint.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("sse-specification", "", "The options for server side encryption.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("tag-specifications", "", "The tags to assign to the Verified Access endpoint.")
	ec2_createVerifiedAccessEndpointCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group to associate the endpoint with.")
	ec2_createVerifiedAccessEndpointCmd.MarkFlagRequired("attachment-type")
	ec2_createVerifiedAccessEndpointCmd.MarkFlagRequired("endpoint-type")
	ec2_createVerifiedAccessEndpointCmd.Flag("no-dry-run").Hidden = true
	ec2_createVerifiedAccessEndpointCmd.MarkFlagRequired("verified-access-group-id")
	ec2Cmd.AddCommand(ec2_createVerifiedAccessEndpointCmd)
}
