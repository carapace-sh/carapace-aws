package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getSipMediaApplicationAlexaSkillConfigurationCmd = &cobra.Command{
	Use:   "get-sip-media-application-alexa-skill-configuration",
	Short: "Gets the Alexa Skill configuration for the SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getSipMediaApplicationAlexaSkillConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getSipMediaApplicationAlexaSkillConfigurationCmd).Standalone()

		chimeSdkVoice_getSipMediaApplicationAlexaSkillConfigurationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
		chimeSdkVoice_getSipMediaApplicationAlexaSkillConfigurationCmd.MarkFlagRequired("sip-media-application-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getSipMediaApplicationAlexaSkillConfigurationCmd)
}
