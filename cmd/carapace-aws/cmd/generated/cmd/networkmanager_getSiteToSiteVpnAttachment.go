package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getSiteToSiteVpnAttachmentCmd = &cobra.Command{
	Use:   "get-site-to-site-vpn-attachment",
	Short: "Returns information about a site-to-site VPN attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getSiteToSiteVpnAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getSiteToSiteVpnAttachmentCmd).Standalone()

		networkmanager_getSiteToSiteVpnAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
		networkmanager_getSiteToSiteVpnAttachmentCmd.MarkFlagRequired("attachment-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getSiteToSiteVpnAttachmentCmd)
}
