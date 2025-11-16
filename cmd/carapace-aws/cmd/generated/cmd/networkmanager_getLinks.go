package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getLinksCmd = &cobra.Command{
	Use:   "get-links",
	Short: "Gets information about one or more links in a specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getLinksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getLinksCmd).Standalone()

		networkmanager_getLinksCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getLinksCmd.Flags().String("link-ids", "", "One or more link IDs.")
		networkmanager_getLinksCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getLinksCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getLinksCmd.Flags().String("provider", "", "The link provider.")
		networkmanager_getLinksCmd.Flags().String("site-id", "", "The ID of the site.")
		networkmanager_getLinksCmd.Flags().String("type", "", "The link type.")
		networkmanager_getLinksCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getLinksCmd)
}
