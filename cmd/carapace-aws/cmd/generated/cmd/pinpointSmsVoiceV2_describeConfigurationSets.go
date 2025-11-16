package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeConfigurationSetsCmd = &cobra.Command{
	Use:   "describe-configuration-sets",
	Short: "Describes the specified configuration sets or all in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeConfigurationSetsCmd).Standalone()

	pinpointSmsVoiceV2_describeConfigurationSetsCmd.Flags().String("configuration-set-names", "", "An array of strings.")
	pinpointSmsVoiceV2_describeConfigurationSetsCmd.Flags().String("filters", "", "An array of filters to apply to the results that are returned.")
	pinpointSmsVoiceV2_describeConfigurationSetsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeConfigurationSetsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeConfigurationSetsCmd)
}
