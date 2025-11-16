package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_batchDeletePhoneNumberCmd = &cobra.Command{
	Use:   "batch-delete-phone-number",
	Short: "Moves phone numbers into the **Deletion queue**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_batchDeletePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_batchDeletePhoneNumberCmd).Standalone()

		chimeSdkVoice_batchDeletePhoneNumberCmd.Flags().String("phone-number-ids", "", "List of phone number IDs.")
		chimeSdkVoice_batchDeletePhoneNumberCmd.MarkFlagRequired("phone-number-ids")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_batchDeletePhoneNumberCmd)
}
