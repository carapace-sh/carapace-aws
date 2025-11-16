package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeProtectConfigurationsCmd = &cobra.Command{
	Use:   "describe-protect-configurations",
	Short: "Retrieves the protect configurations that match any of filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeProtectConfigurationsCmd).Standalone()

	pinpointSmsVoiceV2_describeProtectConfigurationsCmd.Flags().String("filters", "", "An array of ProtectConfigurationFilter objects to filter the results.")
	pinpointSmsVoiceV2_describeProtectConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeProtectConfigurationsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeProtectConfigurationsCmd.Flags().String("protect-configuration-ids", "", "An array of protect configuration identifiers to search for.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeProtectConfigurationsCmd)
}
