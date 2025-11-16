package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteSiteCmd = &cobra.Command{
	Use:   "delete-site",
	Short: "Deletes an existing site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteSiteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteSiteCmd).Standalone()

		networkmanager_deleteSiteCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_deleteSiteCmd.Flags().String("site-id", "", "The ID of the site.")
		networkmanager_deleteSiteCmd.MarkFlagRequired("global-network-id")
		networkmanager_deleteSiteCmd.MarkFlagRequired("site-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteSiteCmd)
}
