package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteIdentityProviderConfigurationCmd = &cobra.Command{
	Use:   "delete-identity-provider-configuration",
	Short: "Disables the integration between IdC and WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteIdentityProviderConfigurationCmd).Standalone()

	workmail_deleteIdentityProviderConfigurationCmd.Flags().String("organization-id", "", "The Organization ID.")
	workmail_deleteIdentityProviderConfigurationCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteIdentityProviderConfigurationCmd)
}
