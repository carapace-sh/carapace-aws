package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessTrustProviderCmd = &cobra.Command{
	Use:   "modify-verified-access-trust-provider",
	Short: "Modifies the configuration of the specified Amazon Web Services Verified Access trust provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessTrustProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVerifiedAccessTrustProviderCmd).Standalone()

		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("description", "", "A description for the Verified Access trust provider.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("device-options", "", "The options for a device-based trust provider.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("native-application-oidc-options", "", "The OpenID Connect (OIDC) options.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("oidc-options", "", "The options for an OpenID Connect-compatible user-identity trust provider.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("sse-specification", "", "The options for server side encryption.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flags().String("verified-access-trust-provider-id", "", "The ID of the Verified Access trust provider.")
		ec2_modifyVerifiedAccessTrustProviderCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVerifiedAccessTrustProviderCmd.MarkFlagRequired("verified-access-trust-provider-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessTrustProviderCmd)
}
