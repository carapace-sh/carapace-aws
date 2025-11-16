package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteDefaultSenderIdCmd = &cobra.Command{
	Use:   "delete-default-sender-id",
	Short: "Deletes an existing default sender ID on a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteDefaultSenderIdCmd).Standalone()

	pinpointSmsVoiceV2_deleteDefaultSenderIdCmd.Flags().String("configuration-set-name", "", "The name of the configuration set or the configuration set Amazon Resource Name (ARN) to delete the default sender ID from.")
	pinpointSmsVoiceV2_deleteDefaultSenderIdCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteDefaultSenderIdCmd)
}
