package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listAvailableVoiceConnectorRegionsCmd = &cobra.Command{
	Use:   "list-available-voice-connector-regions",
	Short: "Lists the available AWS Regions in which you can create an Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listAvailableVoiceConnectorRegionsCmd).Standalone()

	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listAvailableVoiceConnectorRegionsCmd)
}
