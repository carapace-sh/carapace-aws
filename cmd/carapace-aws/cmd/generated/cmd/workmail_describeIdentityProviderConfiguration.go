package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeIdentityProviderConfigurationCmd = &cobra.Command{
	Use:   "describe-identity-provider-configuration",
	Short: "Returns detailed information on the current IdC setup for the WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeIdentityProviderConfigurationCmd).Standalone()

	workmail_describeIdentityProviderConfigurationCmd.Flags().String("organization-id", "", "The Organization ID.")
	workmail_describeIdentityProviderConfigurationCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_describeIdentityProviderConfigurationCmd)
}
