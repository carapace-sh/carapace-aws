package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource-based policy document attached to the End User MessagingSMS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteResourcePolicyCmd).Standalone()

	pinpointSmsVoiceV2_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the End User MessagingSMS resource you're deleting the resource-based policy from.")
	pinpointSmsVoiceV2_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteResourcePolicyCmd)
}
