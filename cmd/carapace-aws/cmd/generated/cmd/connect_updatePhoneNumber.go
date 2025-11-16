package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updatePhoneNumberCmd = &cobra.Command{
	Use:   "update-phone-number",
	Short: "Updates your claimed phone number from its current Amazon Connect instance or traffic distribution group to another Amazon Connect instance or traffic distribution group in the same Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updatePhoneNumberCmd).Standalone()

	connect_updatePhoneNumberCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_updatePhoneNumberCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance that phone numbers are claimed to.")
	connect_updatePhoneNumberCmd.Flags().String("phone-number-id", "", "A unique identifier for the phone number.")
	connect_updatePhoneNumberCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) for Amazon Connect instances or traffic distribution groups that phone number inbound traffic is routed through.")
	connect_updatePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	connectCmd.AddCommand(connect_updatePhoneNumberCmd)
}
