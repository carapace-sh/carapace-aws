package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateEmailAddressAliasCmd = &cobra.Command{
	Use:   "disassociate-email-address-alias",
	Short: "Removes the alias association between two email addresses in an Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateEmailAddressAliasCmd).Standalone()

	connect_disassociateEmailAddressAliasCmd.Flags().String("alias-configuration", "", "Configuration object that specifies which alias relationship to remove.")
	connect_disassociateEmailAddressAliasCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_disassociateEmailAddressAliasCmd.Flags().String("email-address-id", "", "The identifier of the email address.")
	connect_disassociateEmailAddressAliasCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_disassociateEmailAddressAliasCmd.MarkFlagRequired("alias-configuration")
	connect_disassociateEmailAddressAliasCmd.MarkFlagRequired("email-address-id")
	connect_disassociateEmailAddressAliasCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_disassociateEmailAddressAliasCmd)
}
