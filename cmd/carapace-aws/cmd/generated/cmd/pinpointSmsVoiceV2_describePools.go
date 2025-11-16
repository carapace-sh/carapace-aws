package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describePoolsCmd = &cobra.Command{
	Use:   "describe-pools",
	Short: "Retrieves the specified pools or all pools associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describePoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describePoolsCmd).Standalone()

		pinpointSmsVoiceV2_describePoolsCmd.Flags().String("filters", "", "An array of PoolFilter objects to filter the results.")
		pinpointSmsVoiceV2_describePoolsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describePoolsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_describePoolsCmd.Flags().String("owner", "", "Use `SELF` to filter the list of Pools to ones your account owns or use `SHARED` to filter on Pools shared with your account.")
		pinpointSmsVoiceV2_describePoolsCmd.Flags().String("pool-ids", "", "The unique identifier of pools to find.")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describePoolsCmd)
}
