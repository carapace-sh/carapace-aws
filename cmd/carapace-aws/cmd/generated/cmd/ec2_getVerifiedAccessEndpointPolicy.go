package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getVerifiedAccessEndpointPolicyCmd = &cobra.Command{
	Use:   "get-verified-access-endpoint-policy",
	Short: "Get the Verified Access policy associated with the endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getVerifiedAccessEndpointPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getVerifiedAccessEndpointPolicyCmd).Standalone()

		ec2_getVerifiedAccessEndpointPolicyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getVerifiedAccessEndpointPolicyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getVerifiedAccessEndpointPolicyCmd.Flags().String("verified-access-endpoint-id", "", "The ID of the Verified Access endpoint.")
		ec2_getVerifiedAccessEndpointPolicyCmd.Flag("no-dry-run").Hidden = true
		ec2_getVerifiedAccessEndpointPolicyCmd.MarkFlagRequired("verified-access-endpoint-id")
	})
	ec2Cmd.AddCommand(ec2_getVerifiedAccessEndpointPolicyCmd)
}
