package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setDefaultMessageTypeCmd = &cobra.Command{
	Use:   "set-default-message-type",
	Short: "Sets the default message type on a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setDefaultMessageTypeCmd).Standalone()

	pinpointSmsVoiceV2_setDefaultMessageTypeCmd.Flags().String("configuration-set-name", "", "The configuration set to update with a new default message type.")
	pinpointSmsVoiceV2_setDefaultMessageTypeCmd.Flags().String("message-type", "", "The type of message.")
	pinpointSmsVoiceV2_setDefaultMessageTypeCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2_setDefaultMessageTypeCmd.MarkFlagRequired("message-type")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setDefaultMessageTypeCmd)
}
