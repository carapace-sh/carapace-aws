package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getPhoneNumberCmd = &cobra.Command{
	Use:   "get-phone-number",
	Short: "Retrieves details for the specified phone number ID, such as associations, capabilities, and product type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getPhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getPhoneNumberCmd).Standalone()

		chimeSdkVoice_getPhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number ID.")
		chimeSdkVoice_getPhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getPhoneNumberCmd)
}
