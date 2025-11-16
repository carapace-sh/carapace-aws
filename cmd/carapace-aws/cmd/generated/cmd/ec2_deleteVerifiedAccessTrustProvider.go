package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVerifiedAccessTrustProviderCmd = &cobra.Command{
	Use:   "delete-verified-access-trust-provider",
	Short: "Delete an Amazon Web Services Verified Access trust provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVerifiedAccessTrustProviderCmd).Standalone()

	ec2_deleteVerifiedAccessTrustProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_deleteVerifiedAccessTrustProviderCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVerifiedAccessTrustProviderCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVerifiedAccessTrustProviderCmd.Flags().String("verified-access-trust-provider-id", "", "The ID of the Verified Access trust provider.")
	ec2_deleteVerifiedAccessTrustProviderCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteVerifiedAccessTrustProviderCmd.MarkFlagRequired("verified-access-trust-provider-id")
	ec2Cmd.AddCommand(ec2_deleteVerifiedAccessTrustProviderCmd)
}
