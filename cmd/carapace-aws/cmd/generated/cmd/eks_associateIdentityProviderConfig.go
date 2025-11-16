package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_associateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "associate-identity-provider-config",
	Short: "Associates an identity provider configuration to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_associateIdentityProviderConfigCmd).Standalone()

	eks_associateIdentityProviderConfigCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_associateIdentityProviderConfigCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_associateIdentityProviderConfigCmd.Flags().String("oidc", "", "An object representing an OpenID Connect (OIDC) identity provider configuration.")
	eks_associateIdentityProviderConfigCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
	eks_associateIdentityProviderConfigCmd.MarkFlagRequired("cluster-name")
	eks_associateIdentityProviderConfigCmd.MarkFlagRequired("oidc")
	eksCmd.AddCommand(eks_associateIdentityProviderConfigCmd)
}
