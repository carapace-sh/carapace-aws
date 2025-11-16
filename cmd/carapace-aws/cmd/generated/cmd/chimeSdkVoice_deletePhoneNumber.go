package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deletePhoneNumberCmd = &cobra.Command{
	Use:   "delete-phone-number",
	Short: "Moves the specified phone number into the **Deletion queue**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deletePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deletePhoneNumberCmd).Standalone()

		chimeSdkVoice_deletePhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number ID.")
		chimeSdkVoice_deletePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deletePhoneNumberCmd)
}
