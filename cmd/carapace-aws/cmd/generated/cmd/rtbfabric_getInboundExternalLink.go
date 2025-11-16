package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_getInboundExternalLinkCmd = &cobra.Command{
	Use:   "get-inbound-external-link",
	Short: "Retrieves information about an inbound external link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_getInboundExternalLinkCmd).Standalone()

	rtbfabric_getInboundExternalLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_getInboundExternalLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_getInboundExternalLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_getInboundExternalLinkCmd.MarkFlagRequired("link-id")
	rtbfabricCmd.AddCommand(rtbfabric_getInboundExternalLinkCmd)
}
