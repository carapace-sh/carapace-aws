package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd = &cobra.Command{
	Use:   "list-pool-origination-identities",
	Short: "Lists all associated origination identities in your pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd).Standalone()

		pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd.Flags().String("filters", "", "An array of PoolOriginationIdentitiesFilter objects to filter the results..")
		pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd.Flags().String("pool-id", "", "The unique identifier for the pool.")
		pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd.MarkFlagRequired("pool-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_listPoolOriginationIdentitiesCmd)
}
