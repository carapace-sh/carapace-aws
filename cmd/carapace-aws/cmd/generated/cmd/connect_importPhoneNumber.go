package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_importPhoneNumberCmd = &cobra.Command{
	Use:   "import-phone-number",
	Short: "Imports a claimed phone number from an external service, such as Amazon Web Services End User Messaging, into an Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_importPhoneNumberCmd).Standalone()

	connect_importPhoneNumberCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_importPhoneNumberCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_importPhoneNumberCmd.Flags().String("phone-number-description", "", "The description of the phone number.")
	connect_importPhoneNumberCmd.Flags().String("source-phone-number-arn", "", "The claimed phone number ARN being imported from the external service, such as Amazon Web Services End User Messaging.")
	connect_importPhoneNumberCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_importPhoneNumberCmd.MarkFlagRequired("instance-id")
	connect_importPhoneNumberCmd.MarkFlagRequired("source-phone-number-arn")
	connectCmd.AddCommand(connect_importPhoneNumberCmd)
}
