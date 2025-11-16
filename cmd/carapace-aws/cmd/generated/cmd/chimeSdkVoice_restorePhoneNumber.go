package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_restorePhoneNumberCmd = &cobra.Command{
	Use:   "restore-phone-number",
	Short: "Restores a deleted phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_restorePhoneNumberCmd).Standalone()

	chimeSdkVoice_restorePhoneNumberCmd.Flags().String("phone-number-id", "", "The ID of the phone number being restored.")
	chimeSdkVoice_restorePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_restorePhoneNumberCmd)
}
