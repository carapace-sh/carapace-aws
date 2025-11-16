package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_updateLinkCmd = &cobra.Command{
	Use:   "update-link",
	Short: "Updates the configuration of a link between gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_updateLinkCmd).Standalone()

	rtbfabric_updateLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_updateLinkCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_updateLinkCmd.Flags().String("log-settings", "", "Settings for the application logs.")
	rtbfabric_updateLinkCmd.MarkFlagRequired("gateway-id")
	rtbfabric_updateLinkCmd.MarkFlagRequired("link-id")
	rtbfabricCmd.AddCommand(rtbfabric_updateLinkCmd)
}
