package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateEmailAddressAliasCmd = &cobra.Command{
	Use:   "associate-email-address-alias",
	Short: "Associates an email address alias with an existing email address in an Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateEmailAddressAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_associateEmailAddressAliasCmd).Standalone()

		connect_associateEmailAddressAliasCmd.Flags().String("alias-configuration", "", "Configuration object that specifies which email address will serve as the alias.")
		connect_associateEmailAddressAliasCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_associateEmailAddressAliasCmd.Flags().String("email-address-id", "", "The identifier of the email address.")
		connect_associateEmailAddressAliasCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_associateEmailAddressAliasCmd.MarkFlagRequired("alias-configuration")
		connect_associateEmailAddressAliasCmd.MarkFlagRequired("email-address-id")
		connect_associateEmailAddressAliasCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_associateEmailAddressAliasCmd)
}
