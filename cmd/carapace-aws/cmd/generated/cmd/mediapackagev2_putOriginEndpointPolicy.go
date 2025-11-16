package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_putOriginEndpointPolicyCmd = &cobra.Command{
	Use:   "put-origin-endpoint-policy",
	Short: "Attaches an IAM policy to the specified origin endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_putOriginEndpointPolicyCmd).Standalone()

	mediapackagev2_putOriginEndpointPolicyCmd.Flags().String("cdn-auth-configuration", "", "The settings for using authorization headers between the MediaPackage endpoint and your CDN.")
	mediapackagev2_putOriginEndpointPolicyCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_putOriginEndpointPolicyCmd.Flags().String("channel-name", "", "The name that describes the channel.")
	mediapackagev2_putOriginEndpointPolicyCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
	mediapackagev2_putOriginEndpointPolicyCmd.Flags().String("policy", "", "The policy to attach to the specified origin endpoint.")
	mediapackagev2_putOriginEndpointPolicyCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_putOriginEndpointPolicyCmd.MarkFlagRequired("channel-name")
	mediapackagev2_putOriginEndpointPolicyCmd.MarkFlagRequired("origin-endpoint-name")
	mediapackagev2_putOriginEndpointPolicyCmd.MarkFlagRequired("policy")
	mediapackagev2Cmd.AddCommand(mediapackagev2_putOriginEndpointPolicyCmd)
}
