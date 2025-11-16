package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_getOutboundExternalLinkCmd = &cobra.Command{
	Use:   "get-outbound-external-link",
	Short: "Retrieves information about an outbound external link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_getOutboundExternalLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_getOutboundExternalLinkCmd).Standalone()

		rtbfabric_getOutboundExternalLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_getOutboundExternalLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
		rtbfabric_getOutboundExternalLinkCmd.MarkFlagRequired("gateway-id")
		rtbfabric_getOutboundExternalLinkCmd.MarkFlagRequired("link-id")
	})
	rtbfabricCmd.AddCommand(rtbfabric_getOutboundExternalLinkCmd)
}
