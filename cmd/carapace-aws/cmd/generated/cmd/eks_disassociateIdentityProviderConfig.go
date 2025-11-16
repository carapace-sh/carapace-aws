package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_disassociateIdentityProviderConfigCmd = &cobra.Command{
	Use:   "disassociate-identity-provider-config",
	Short: "Disassociates an identity provider configuration from a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_disassociateIdentityProviderConfigCmd).Standalone()

	eks_disassociateIdentityProviderConfigCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_disassociateIdentityProviderConfigCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_disassociateIdentityProviderConfigCmd.Flags().String("identity-provider-config", "", "An object representing an identity provider configuration.")
	eks_disassociateIdentityProviderConfigCmd.MarkFlagRequired("cluster-name")
	eks_disassociateIdentityProviderConfigCmd.MarkFlagRequired("identity-provider-config")
	eksCmd.AddCommand(eks_disassociateIdentityProviderConfigCmd)
}
