package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_listConfigurationSetsCmd = &cobra.Command{
	Use:   "list-configuration-sets",
	Short: "List all of the configuration sets associated with your Amazon Pinpoint account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_listConfigurationSetsCmd).Standalone()

	pinpointSmsVoice_listConfigurationSetsCmd.Flags().String("next-token", "", "A token returned from a previous call to the API that indicates the position in the list of results.")
	pinpointSmsVoice_listConfigurationSetsCmd.Flags().String("page-size", "", "Used to specify the number of items that should be returned in the response.")
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_listConfigurationSetsCmd)
}
