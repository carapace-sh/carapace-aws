package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessEndpointPolicyCmd = &cobra.Command{
	Use:   "modify-verified-access-endpoint-policy",
	Short: "Modifies the specified Amazon Web Services Verified Access endpoint policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessEndpointPolicyCmd).Standalone()

	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().Bool("no-policy-enabled", false, "The status of the Verified Access policy.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().String("policy-document", "", "The Verified Access policy document.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().Bool("policy-enabled", false, "The status of the Verified Access policy.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().String("sse-specification", "", "The options for server side encryption.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flags().String("verified-access-endpoint-id", "", "The ID of the Verified Access endpoint.")
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVerifiedAccessEndpointPolicyCmd.Flag("no-policy-enabled").Hidden = true
	ec2_modifyVerifiedAccessEndpointPolicyCmd.MarkFlagRequired("verified-access-endpoint-id")
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessEndpointPolicyCmd)
}
