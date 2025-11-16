package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_deleteInboundExternalLinkCmd = &cobra.Command{
	Use:   "delete-inbound-external-link",
	Short: "Deletes an inbound external link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_deleteInboundExternalLinkCmd).Standalone()

	rtbfabric_deleteInboundExternalLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_deleteInboundExternalLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_deleteInboundExternalLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_deleteInboundExternalLinkCmd.MarkFlagRequired("link-id")
	rtbfabricCmd.AddCommand(rtbfabric_deleteInboundExternalLinkCmd)
}
