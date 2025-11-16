package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getSitesCmd = &cobra.Command{
	Use:   "get-sites",
	Short: "Gets information about one or more of your sites in a global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getSitesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getSitesCmd).Standalone()

		networkmanager_getSitesCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getSitesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getSitesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getSitesCmd.Flags().String("site-ids", "", "One or more site IDs.")
		networkmanager_getSitesCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getSitesCmd)
}
