package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_attachVerifiedAccessTrustProviderCmd = &cobra.Command{
	Use:   "attach-verified-access-trust-provider",
	Short: "Attaches the specified Amazon Web Services Verified Access trust provider to the specified Amazon Web Services Verified Access instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_attachVerifiedAccessTrustProviderCmd).Standalone()

	ec2_attachVerifiedAccessTrustProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_attachVerifiedAccessTrustProviderCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_attachVerifiedAccessTrustProviderCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_attachVerifiedAccessTrustProviderCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
	ec2_attachVerifiedAccessTrustProviderCmd.Flags().String("verified-access-trust-provider-id", "", "The ID of the Verified Access trust provider.")
	ec2_attachVerifiedAccessTrustProviderCmd.Flag("no-dry-run").Hidden = true
	ec2_attachVerifiedAccessTrustProviderCmd.MarkFlagRequired("verified-access-instance-id")
	ec2_attachVerifiedAccessTrustProviderCmd.MarkFlagRequired("verified-access-trust-provider-id")
	ec2Cmd.AddCommand(ec2_attachVerifiedAccessTrustProviderCmd)
}
