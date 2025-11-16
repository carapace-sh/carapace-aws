package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updatePhoneNumberMetadataCmd = &cobra.Command{
	Use:   "update-phone-number-metadata",
	Short: "Updates a phone number’s metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updatePhoneNumberMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updatePhoneNumberMetadataCmd).Standalone()

		connect_updatePhoneNumberMetadataCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_updatePhoneNumberMetadataCmd.Flags().String("phone-number-description", "", "The description of the phone number.")
		connect_updatePhoneNumberMetadataCmd.Flags().String("phone-number-id", "", "The Amazon Resource Name (ARN) or resource ID of the phone number.")
		connect_updatePhoneNumberMetadataCmd.MarkFlagRequired("phone-number-id")
	})
	connectCmd.AddCommand(connect_updatePhoneNumberMetadataCmd)
}
