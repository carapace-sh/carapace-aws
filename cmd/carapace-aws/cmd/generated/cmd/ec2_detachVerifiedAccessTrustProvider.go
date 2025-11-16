package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_detachVerifiedAccessTrustProviderCmd = &cobra.Command{
	Use:   "detach-verified-access-trust-provider",
	Short: "Detaches the specified Amazon Web Services Verified Access trust provider from the specified Amazon Web Services Verified Access instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_detachVerifiedAccessTrustProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_detachVerifiedAccessTrustProviderCmd).Standalone()

		ec2_detachVerifiedAccessTrustProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_detachVerifiedAccessTrustProviderCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachVerifiedAccessTrustProviderCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachVerifiedAccessTrustProviderCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_detachVerifiedAccessTrustProviderCmd.Flags().String("verified-access-trust-provider-id", "", "The ID of the Verified Access trust provider.")
		ec2_detachVerifiedAccessTrustProviderCmd.Flag("no-dry-run").Hidden = true
		ec2_detachVerifiedAccessTrustProviderCmd.MarkFlagRequired("verified-access-instance-id")
		ec2_detachVerifiedAccessTrustProviderCmd.MarkFlagRequired("verified-access-trust-provider-id")
	})
	ec2Cmd.AddCommand(ec2_detachVerifiedAccessTrustProviderCmd)
}
