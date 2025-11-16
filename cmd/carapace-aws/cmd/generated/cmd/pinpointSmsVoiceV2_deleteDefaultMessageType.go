package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteDefaultMessageTypeCmd = &cobra.Command{
	Use:   "delete-default-message-type",
	Short: "Deletes an existing default message type on a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteDefaultMessageTypeCmd).Standalone()

	pinpointSmsVoiceV2_deleteDefaultMessageTypeCmd.Flags().String("configuration-set-name", "", "The name of the configuration set or the configuration set Amazon Resource Name (ARN) to delete the default message type from.")
	pinpointSmsVoiceV2_deleteDefaultMessageTypeCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteDefaultMessageTypeCmd)
}
