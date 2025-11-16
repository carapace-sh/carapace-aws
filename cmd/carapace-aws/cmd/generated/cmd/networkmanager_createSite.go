package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createSiteCmd = &cobra.Command{
	Use:   "create-site",
	Short: "Creates a new site in a global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createSiteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createSiteCmd).Standalone()

		networkmanager_createSiteCmd.Flags().String("description", "", "A description of your site.")
		networkmanager_createSiteCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_createSiteCmd.Flags().String("location", "", "The site location.")
		networkmanager_createSiteCmd.Flags().String("tags", "", "The tags to apply to the resource during creation.")
		networkmanager_createSiteCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_createSiteCmd)
}
