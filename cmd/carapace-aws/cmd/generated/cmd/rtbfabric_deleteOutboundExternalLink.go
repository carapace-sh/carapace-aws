package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_deleteOutboundExternalLinkCmd = &cobra.Command{
	Use:   "delete-outbound-external-link",
	Short: "Deletes an outbound external link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_deleteOutboundExternalLinkCmd).Standalone()

	rtbfabric_deleteOutboundExternalLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_deleteOutboundExternalLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_deleteOutboundExternalLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_deleteOutboundExternalLinkCmd.MarkFlagRequired("link-id")
	rtbfabricCmd.AddCommand(rtbfabric_deleteOutboundExternalLinkCmd)
}
