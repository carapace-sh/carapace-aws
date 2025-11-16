package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createSiteToSiteVpnAttachmentCmd = &cobra.Command{
	Use:   "create-site-to-site-vpn-attachment",
	Short: "Creates an Amazon Web Services site-to-site VPN attachment on an edge location of a core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createSiteToSiteVpnAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createSiteToSiteVpnAttachmentCmd).Standalone()

		networkmanager_createSiteToSiteVpnAttachmentCmd.Flags().String("client-token", "", "The client token associated with the request.")
		networkmanager_createSiteToSiteVpnAttachmentCmd.Flags().String("core-network-id", "", "The ID of a core network where you're creating a site-to-site VPN attachment.")
		networkmanager_createSiteToSiteVpnAttachmentCmd.Flags().String("tags", "", "The tags associated with the request.")
		networkmanager_createSiteToSiteVpnAttachmentCmd.Flags().String("vpn-connection-arn", "", "The ARN identifying the VPN attachment.")
		networkmanager_createSiteToSiteVpnAttachmentCmd.MarkFlagRequired("core-network-id")
		networkmanager_createSiteToSiteVpnAttachmentCmd.MarkFlagRequired("vpn-connection-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_createSiteToSiteVpnAttachmentCmd)
}
