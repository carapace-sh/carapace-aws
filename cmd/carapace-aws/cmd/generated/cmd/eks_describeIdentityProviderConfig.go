package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeIdentityProviderConfigCmd = &cobra.Command{
	Use:   "describe-identity-provider-config",
	Short: "Describes an identity provider configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeIdentityProviderConfigCmd).Standalone()

	eks_describeIdentityProviderConfigCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_describeIdentityProviderConfigCmd.Flags().String("identity-provider-config", "", "An object representing an identity provider configuration.")
	eks_describeIdentityProviderConfigCmd.MarkFlagRequired("cluster-name")
	eks_describeIdentityProviderConfigCmd.MarkFlagRequired("identity-provider-config")
	eksCmd.AddCommand(eks_describeIdentityProviderConfigCmd)
}
