package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_createOutboundExternalLinkCmd = &cobra.Command{
	Use:   "create-outbound-external-link",
	Short: "Creates an outbound external link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_createOutboundExternalLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_createOutboundExternalLinkCmd).Standalone()

		rtbfabric_createOutboundExternalLinkCmd.Flags().String("attributes", "", "")
		rtbfabric_createOutboundExternalLinkCmd.Flags().String("client-token", "", "The unique client token.")
		rtbfabric_createOutboundExternalLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_createOutboundExternalLinkCmd.Flags().String("log-settings", "", "")
		rtbfabric_createOutboundExternalLinkCmd.Flags().String("public-endpoint", "", "The public endpoint of the link.")
		rtbfabric_createOutboundExternalLinkCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
		rtbfabric_createOutboundExternalLinkCmd.MarkFlagRequired("client-token")
		rtbfabric_createOutboundExternalLinkCmd.MarkFlagRequired("gateway-id")
		rtbfabric_createOutboundExternalLinkCmd.MarkFlagRequired("log-settings")
		rtbfabric_createOutboundExternalLinkCmd.MarkFlagRequired("public-endpoint")
	})
	rtbfabricCmd.AddCommand(rtbfabric_createOutboundExternalLinkCmd)
}
