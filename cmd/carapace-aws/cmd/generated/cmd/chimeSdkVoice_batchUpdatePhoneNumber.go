package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_batchUpdatePhoneNumberCmd = &cobra.Command{
	Use:   "batch-update-phone-number",
	Short: "Updates phone number product types, calling names, or phone number names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_batchUpdatePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_batchUpdatePhoneNumberCmd).Standalone()

		chimeSdkVoice_batchUpdatePhoneNumberCmd.Flags().String("update-phone-number-request-items", "", "Lists the phone numbers in the update request.")
		chimeSdkVoice_batchUpdatePhoneNumberCmd.MarkFlagRequired("update-phone-number-request-items")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_batchUpdatePhoneNumberCmd)
}
