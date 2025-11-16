package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_createInboundExternalLinkCmd = &cobra.Command{
	Use:   "create-inbound-external-link",
	Short: "Creates an inbound external link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_createInboundExternalLinkCmd).Standalone()

	rtbfabric_createInboundExternalLinkCmd.Flags().String("attributes", "", "Attributes of the link.")
	rtbfabric_createInboundExternalLinkCmd.Flags().String("client-token", "", "The unique client token.")
	rtbfabric_createInboundExternalLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_createInboundExternalLinkCmd.Flags().String("log-settings", "", "")
	rtbfabric_createInboundExternalLinkCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	rtbfabric_createInboundExternalLinkCmd.MarkFlagRequired("client-token")
	rtbfabric_createInboundExternalLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_createInboundExternalLinkCmd.MarkFlagRequired("log-settings")
	rtbfabricCmd.AddCommand(rtbfabric_createInboundExternalLinkCmd)
}
