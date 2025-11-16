package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the JSON text of the resource-based policy document attached to the End User MessagingSMS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_getResourcePolicyCmd).Standalone()

	pinpointSmsVoiceV2_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the End User MessagingSMS resource attached to the resource-based policy.")
	pinpointSmsVoiceV2_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_getResourcePolicyCmd)
}
