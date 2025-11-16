package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateLinkCmd = &cobra.Command{
	Use:   "update-link",
	Short: "Updates the details for an existing link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_updateLinkCmd).Standalone()

		networkmanager_updateLinkCmd.Flags().String("bandwidth", "", "The upload and download speed in Mbps.")
		networkmanager_updateLinkCmd.Flags().String("description", "", "A description of the link.")
		networkmanager_updateLinkCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_updateLinkCmd.Flags().String("link-id", "", "The ID of the link.")
		networkmanager_updateLinkCmd.Flags().String("provider", "", "The provider of the link.")
		networkmanager_updateLinkCmd.Flags().String("type", "", "The type of the link.")
		networkmanager_updateLinkCmd.MarkFlagRequired("global-network-id")
		networkmanager_updateLinkCmd.MarkFlagRequired("link-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_updateLinkCmd)
}
