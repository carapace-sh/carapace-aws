package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateGlobalSettingsCmd = &cobra.Command{
	Use:   "update-global-settings",
	Short: "Updates global settings for the administrator's AWS account, such as Amazon Chime Business Calling and Amazon Chime Voice Connector settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateGlobalSettingsCmd).Standalone()

	chime_updateGlobalSettingsCmd.Flags().String("business-calling", "", "The Amazon Chime Business Calling settings.")
	chime_updateGlobalSettingsCmd.Flags().String("voice-connector", "", "The Amazon Chime Voice Connector settings.")
	chimeCmd.AddCommand(chime_updateGlobalSettingsCmd)
}
