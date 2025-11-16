package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_deleteOriginEndpointPolicyCmd = &cobra.Command{
	Use:   "delete-origin-endpoint-policy",
	Short: "Delete an origin endpoint policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_deleteOriginEndpointPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_deleteOriginEndpointPolicyCmd).Standalone()

		mediapackagev2_deleteOriginEndpointPolicyCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_deleteOriginEndpointPolicyCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_deleteOriginEndpointPolicyCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
		mediapackagev2_deleteOriginEndpointPolicyCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_deleteOriginEndpointPolicyCmd.MarkFlagRequired("channel-name")
		mediapackagev2_deleteOriginEndpointPolicyCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_deleteOriginEndpointPolicyCmd)
}
