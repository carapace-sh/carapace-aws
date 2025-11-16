package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd = &cobra.Command{
	Use:   "set-default-message-feedback-enabled",
	Short: "Sets a configuration set's default for message feedback.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd).Standalone()

		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use.")
		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.Flags().Bool("message-feedback-enabled", false, "Set to true to enable message feedback.")
		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.Flags().Bool("no-message-feedback-enabled", false, "Set to true to enable message feedback.")
		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.MarkFlagRequired("configuration-set-name")
		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.MarkFlagRequired("message-feedback-enabled")
		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.Flag("no-message-feedback-enabled").Hidden = true
		pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd.MarkFlagRequired("no-message-feedback-enabled")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setDefaultMessageFeedbackEnabledCmd)
}
