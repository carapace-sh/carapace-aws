package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_claimPhoneNumberCmd = &cobra.Command{
	Use:   "claim-phone-number",
	Short: "Claims an available phone number to your Amazon Connect instance or traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_claimPhoneNumberCmd).Standalone()

	connect_claimPhoneNumberCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_claimPhoneNumberCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance that phone numbers are claimed to.")
	connect_claimPhoneNumberCmd.Flags().String("phone-number", "", "The phone number you want to claim.")
	connect_claimPhoneNumberCmd.Flags().String("phone-number-description", "", "The description of the phone number.")
	connect_claimPhoneNumberCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_claimPhoneNumberCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) for Amazon Connect instances or traffic distribution groups that phone number inbound traffic is routed through.")
	connect_claimPhoneNumberCmd.MarkFlagRequired("phone-number")
	connectCmd.AddCommand(connect_claimPhoneNumberCmd)
}
