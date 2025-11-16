package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVerifiedAccessTrustProviderCmd = &cobra.Command{
	Use:   "create-verified-access-trust-provider",
	Short: "A trust provider is a third-party entity that creates, maintains, and manages identity information for users and devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVerifiedAccessTrustProviderCmd).Standalone()

	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("description", "", "A description for the Verified Access trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("device-options", "", "The options for a device-based trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("device-trust-provider-type", "", "The type of device-based trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("native-application-oidc-options", "", "The OpenID Connect (OIDC) options.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("oidc-options", "", "The options for a OpenID Connect-compatible user-identity trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("policy-reference-name", "", "The identifier to be used when working with policy rules.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("sse-specification", "", "The options for server side encryption.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("tag-specifications", "", "The tags to assign to the Verified Access trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("trust-provider-type", "", "The type of trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flags().String("user-trust-provider-type", "", "The type of user-based trust provider.")
	ec2_createVerifiedAccessTrustProviderCmd.Flag("no-dry-run").Hidden = true
	ec2_createVerifiedAccessTrustProviderCmd.MarkFlagRequired("policy-reference-name")
	ec2_createVerifiedAccessTrustProviderCmd.MarkFlagRequired("trust-provider-type")
	ec2Cmd.AddCommand(ec2_createVerifiedAccessTrustProviderCmd)
}
