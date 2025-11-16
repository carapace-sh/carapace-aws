package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putSipMediaApplicationAlexaSkillConfigurationCmd = &cobra.Command{
	Use:   "put-sip-media-application-alexa-skill-configuration",
	Short: "Updates the Alexa Skill configuration for the SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putSipMediaApplicationAlexaSkillConfigurationCmd).Standalone()

	chimeSdkVoice_putSipMediaApplicationAlexaSkillConfigurationCmd.Flags().String("sip-media-application-alexa-skill-configuration", "", "The Alexa Skill configuration.")
	chimeSdkVoice_putSipMediaApplicationAlexaSkillConfigurationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
	chimeSdkVoice_putSipMediaApplicationAlexaSkillConfigurationCmd.MarkFlagRequired("sip-media-application-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putSipMediaApplicationAlexaSkillConfigurationCmd)
}
