package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_disassociateLinkCmd = &cobra.Command{
	Use:   "disassociate-link",
	Short: "Disassociates an existing device from a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_disassociateLinkCmd).Standalone()

	networkmanager_disassociateLinkCmd.Flags().String("device-id", "", "The ID of the device.")
	networkmanager_disassociateLinkCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_disassociateLinkCmd.Flags().String("link-id", "", "The ID of the link.")
	networkmanager_disassociateLinkCmd.MarkFlagRequired("device-id")
	networkmanager_disassociateLinkCmd.MarkFlagRequired("global-network-id")
	networkmanager_disassociateLinkCmd.MarkFlagRequired("link-id")
	networkmanagerCmd.AddCommand(networkmanager_disassociateLinkCmd)
}
