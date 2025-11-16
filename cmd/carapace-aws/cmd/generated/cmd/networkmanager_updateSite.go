package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateSiteCmd = &cobra.Command{
	Use:   "update-site",
	Short: "Updates the information for an existing site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateSiteCmd).Standalone()

	networkmanager_updateSiteCmd.Flags().String("description", "", "A description of your site.")
	networkmanager_updateSiteCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_updateSiteCmd.Flags().String("location", "", "The site location:")
	networkmanager_updateSiteCmd.Flags().String("site-id", "", "The ID of your site.")
	networkmanager_updateSiteCmd.MarkFlagRequired("global-network-id")
	networkmanager_updateSiteCmd.MarkFlagRequired("site-id")
	networkmanagerCmd.AddCommand(networkmanager_updateSiteCmd)
}
