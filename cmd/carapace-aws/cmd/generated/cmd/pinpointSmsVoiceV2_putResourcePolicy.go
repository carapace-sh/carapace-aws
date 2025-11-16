package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based policy to a End User MessagingSMS resource(phone number, sender Id, phone poll, or opt-out list) that is used for sharing the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_putResourcePolicyCmd).Standalone()

		pinpointSmsVoiceV2_putResourcePolicyCmd.Flags().String("policy", "", "The JSON formatted resource-based policy to attach.")
		pinpointSmsVoiceV2_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the End User MessagingSMS resource to attach the resource-based policy to.")
		pinpointSmsVoiceV2_putResourcePolicyCmd.MarkFlagRequired("policy")
		pinpointSmsVoiceV2_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_putResourcePolicyCmd)
}
