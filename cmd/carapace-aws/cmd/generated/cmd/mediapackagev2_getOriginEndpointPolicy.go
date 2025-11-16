package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_getOriginEndpointPolicyCmd = &cobra.Command{
	Use:   "get-origin-endpoint-policy",
	Short: "Retrieves the specified origin endpoint policy that's configured in AWS Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_getOriginEndpointPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_getOriginEndpointPolicyCmd).Standalone()

		mediapackagev2_getOriginEndpointPolicyCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_getOriginEndpointPolicyCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_getOriginEndpointPolicyCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
		mediapackagev2_getOriginEndpointPolicyCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_getOriginEndpointPolicyCmd.MarkFlagRequired("channel-name")
		mediapackagev2_getOriginEndpointPolicyCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_getOriginEndpointPolicyCmd)
}
