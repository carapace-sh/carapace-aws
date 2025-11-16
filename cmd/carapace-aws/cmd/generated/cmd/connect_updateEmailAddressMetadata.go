package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateEmailAddressMetadataCmd = &cobra.Command{
	Use:   "update-email-address-metadata",
	Short: "Updates an email address metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateEmailAddressMetadataCmd).Standalone()

	connect_updateEmailAddressMetadataCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_updateEmailAddressMetadataCmd.Flags().String("description", "", "The description of the email address.")
	connect_updateEmailAddressMetadataCmd.Flags().String("display-name", "", "The display name of email address.")
	connect_updateEmailAddressMetadataCmd.Flags().String("email-address-id", "", "The identifier of the email address.")
	connect_updateEmailAddressMetadataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateEmailAddressMetadataCmd.MarkFlagRequired("email-address-id")
	connect_updateEmailAddressMetadataCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateEmailAddressMetadataCmd)
}
