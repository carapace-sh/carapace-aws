package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_associateLinkCmd = &cobra.Command{
	Use:   "associate-link",
	Short: "Associates a link to a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_associateLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_associateLinkCmd).Standalone()

		networkmanager_associateLinkCmd.Flags().String("device-id", "", "The ID of the device.")
		networkmanager_associateLinkCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_associateLinkCmd.Flags().String("link-id", "", "The ID of the link.")
		networkmanager_associateLinkCmd.MarkFlagRequired("device-id")
		networkmanager_associateLinkCmd.MarkFlagRequired("global-network-id")
		networkmanager_associateLinkCmd.MarkFlagRequired("link-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_associateLinkCmd)
}
