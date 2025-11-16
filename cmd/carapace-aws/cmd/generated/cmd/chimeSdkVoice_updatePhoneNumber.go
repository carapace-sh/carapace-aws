package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updatePhoneNumberCmd = &cobra.Command{
	Use:   "update-phone-number",
	Short: "Updates phone number details, such as product type, calling name, or phone number name for the specified phone number ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updatePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_updatePhoneNumberCmd).Standalone()

		chimeSdkVoice_updatePhoneNumberCmd.Flags().String("calling-name", "", "The outbound calling name associated with the phone number.")
		chimeSdkVoice_updatePhoneNumberCmd.Flags().String("name", "", "Specifies the updated name assigned to one or more phone numbers.")
		chimeSdkVoice_updatePhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number ID.")
		chimeSdkVoice_updatePhoneNumberCmd.Flags().String("product-type", "", "The product type.")
		chimeSdkVoice_updatePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updatePhoneNumberCmd)
}
