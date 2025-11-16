package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteKeywordCmd = &cobra.Command{
	Use:   "delete-keyword",
	Short: "Deletes an existing keyword from an origination phone number or pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteKeywordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_deleteKeywordCmd).Standalone()

		pinpointSmsVoiceV2_deleteKeywordCmd.Flags().String("keyword", "", "The keyword to delete.")
		pinpointSmsVoiceV2_deleteKeywordCmd.Flags().String("origination-identity", "", "The origination identity to use such as a PhoneNumberId, PhoneNumberArn, PoolId or PoolArn.")
		pinpointSmsVoiceV2_deleteKeywordCmd.MarkFlagRequired("keyword")
		pinpointSmsVoiceV2_deleteKeywordCmd.MarkFlagRequired("origination-identity")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteKeywordCmd)
}
