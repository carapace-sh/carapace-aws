package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putIdentityProviderConfigurationCmd = &cobra.Command{
	Use:   "put-identity-provider-configuration",
	Short: "Enables integration between IAM Identity Center (IdC) and WorkMail to proxy authentication requests for mailbox users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putIdentityProviderConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_putIdentityProviderConfigurationCmd).Standalone()

		workmail_putIdentityProviderConfigurationCmd.Flags().String("authentication-mode", "", "The authentication mode used in WorkMail.")
		workmail_putIdentityProviderConfigurationCmd.Flags().String("identity-center-configuration", "", "The details of the IAM Identity Center configuration.")
		workmail_putIdentityProviderConfigurationCmd.Flags().String("organization-id", "", "The ID of the WorkMail Organization.")
		workmail_putIdentityProviderConfigurationCmd.Flags().String("personal-access-token-configuration", "", "The details of the Personal Access Token configuration.")
		workmail_putIdentityProviderConfigurationCmd.MarkFlagRequired("authentication-mode")
		workmail_putIdentityProviderConfigurationCmd.MarkFlagRequired("identity-center-configuration")
		workmail_putIdentityProviderConfigurationCmd.MarkFlagRequired("organization-id")
		workmail_putIdentityProviderConfigurationCmd.MarkFlagRequired("personal-access-token-configuration")
	})
	workmailCmd.AddCommand(workmail_putIdentityProviderConfigurationCmd)
}
