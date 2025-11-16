package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_deleteLinkCmd = &cobra.Command{
	Use:   "delete-link",
	Short: "Deletes a link between gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_deleteLinkCmd).Standalone()

	rtbfabric_deleteLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_deleteLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_deleteLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_deleteLinkCmd.MarkFlagRequired("link-id")
	rtbfabricCmd.AddCommand(rtbfabric_deleteLinkCmd)
}
