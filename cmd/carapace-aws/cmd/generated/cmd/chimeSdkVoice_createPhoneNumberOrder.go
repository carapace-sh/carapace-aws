package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createPhoneNumberOrderCmd = &cobra.Command{
	Use:   "create-phone-number-order",
	Short: "Creates an order for phone numbers to be provisioned.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createPhoneNumberOrderCmd).Standalone()

	chimeSdkVoice_createPhoneNumberOrderCmd.Flags().String("e164-phone-numbers", "", "List of phone numbers, in E.164 format.")
	chimeSdkVoice_createPhoneNumberOrderCmd.Flags().String("name", "", "Specifies the name assigned to one or more phone numbers.")
	chimeSdkVoice_createPhoneNumberOrderCmd.Flags().String("product-type", "", "The phone number product type.")
	chimeSdkVoice_createPhoneNumberOrderCmd.MarkFlagRequired("e164-phone-numbers")
	chimeSdkVoice_createPhoneNumberOrderCmd.MarkFlagRequired("product-type")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createPhoneNumberOrderCmd)
}
