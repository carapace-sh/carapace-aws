package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_putKeywordCmd = &cobra.Command{
	Use:   "put-keyword",
	Short: "Creates or updates a keyword configuration on an origination phone number or pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_putKeywordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_putKeywordCmd).Standalone()

		pinpointSmsVoiceV2_putKeywordCmd.Flags().String("keyword", "", "The new keyword to add.")
		pinpointSmsVoiceV2_putKeywordCmd.Flags().String("keyword-action", "", "The action to perform for the new keyword when it is received.")
		pinpointSmsVoiceV2_putKeywordCmd.Flags().String("keyword-message", "", "The message associated with the keyword.")
		pinpointSmsVoiceV2_putKeywordCmd.Flags().String("origination-identity", "", "The origination identity to use such as a PhoneNumberId, PhoneNumberArn, SenderId or SenderIdArn.")
		pinpointSmsVoiceV2_putKeywordCmd.MarkFlagRequired("keyword")
		pinpointSmsVoiceV2_putKeywordCmd.MarkFlagRequired("keyword-message")
		pinpointSmsVoiceV2_putKeywordCmd.MarkFlagRequired("origination-identity")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_putKeywordCmd)
}
