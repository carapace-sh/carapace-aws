package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_getLinkCmd = &cobra.Command{
	Use:   "get-link",
	Short: "Retrieves information about a link between gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_getLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_getLinkCmd).Standalone()

		rtbfabric_getLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_getLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
		rtbfabric_getLinkCmd.MarkFlagRequired("gateway-id")
		rtbfabric_getLinkCmd.MarkFlagRequired("link-id")
	})
	rtbfabricCmd.AddCommand(rtbfabric_getLinkCmd)
}
