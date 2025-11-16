package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createLinkCmd = &cobra.Command{
	Use:   "create-link",
	Short: "Creates a new link for a specified site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createLinkCmd).Standalone()

		networkmanager_createLinkCmd.Flags().String("bandwidth", "", "The upload speed and download speed in Mbps.")
		networkmanager_createLinkCmd.Flags().String("description", "", "A description of the link.")
		networkmanager_createLinkCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_createLinkCmd.Flags().String("provider", "", "The provider of the link.")
		networkmanager_createLinkCmd.Flags().String("site-id", "", "The ID of the site.")
		networkmanager_createLinkCmd.Flags().String("tags", "", "The tags to apply to the resource during creation.")
		networkmanager_createLinkCmd.Flags().String("type", "", "The type of the link.")
		networkmanager_createLinkCmd.MarkFlagRequired("bandwidth")
		networkmanager_createLinkCmd.MarkFlagRequired("global-network-id")
		networkmanager_createLinkCmd.MarkFlagRequired("site-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_createLinkCmd)
}
