package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_acceptLinkCmd = &cobra.Command{
	Use:   "accept-link",
	Short: "Accepts a link request between gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_acceptLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_acceptLinkCmd).Standalone()

		rtbfabric_acceptLinkCmd.Flags().String("attributes", "", "Attributes of the link.")
		rtbfabric_acceptLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_acceptLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
		rtbfabric_acceptLinkCmd.Flags().String("log-settings", "", "Settings for the application logs.")
		rtbfabric_acceptLinkCmd.MarkFlagRequired("gateway-id")
		rtbfabric_acceptLinkCmd.MarkFlagRequired("link-id")
		rtbfabric_acceptLinkCmd.MarkFlagRequired("log-settings")
	})
	rtbfabricCmd.AddCommand(rtbfabric_acceptLinkCmd)
}
