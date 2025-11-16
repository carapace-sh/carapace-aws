package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setDefaultSenderIdCmd = &cobra.Command{
	Use:   "set-default-sender-id",
	Short: "Sets default sender ID on a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setDefaultSenderIdCmd).Standalone()

	pinpointSmsVoiceV2_setDefaultSenderIdCmd.Flags().String("configuration-set-name", "", "The configuration set to updated with a new default SenderId.")
	pinpointSmsVoiceV2_setDefaultSenderIdCmd.Flags().String("sender-id", "", "The current sender ID for the configuration set.")
	pinpointSmsVoiceV2_setDefaultSenderIdCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2_setDefaultSenderIdCmd.MarkFlagRequired("sender-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setDefaultSenderIdCmd)
}
