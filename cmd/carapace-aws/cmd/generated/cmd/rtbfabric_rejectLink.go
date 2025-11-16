package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_rejectLinkCmd = &cobra.Command{
	Use:   "reject-link",
	Short: "Rejects a link request between gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_rejectLinkCmd).Standalone()

	rtbfabric_rejectLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_rejectLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_rejectLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_rejectLinkCmd.MarkFlagRequired("link-id")
	rtbfabricCmd.AddCommand(rtbfabric_rejectLinkCmd)
}
