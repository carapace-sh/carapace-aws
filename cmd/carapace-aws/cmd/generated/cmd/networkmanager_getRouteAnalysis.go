package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getRouteAnalysisCmd = &cobra.Command{
	Use:   "get-route-analysis",
	Short: "Gets information about the specified route analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getRouteAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getRouteAnalysisCmd).Standalone()

		networkmanager_getRouteAnalysisCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getRouteAnalysisCmd.Flags().String("route-analysis-id", "", "The ID of the route analysis.")
		networkmanager_getRouteAnalysisCmd.MarkFlagRequired("global-network-id")
		networkmanager_getRouteAnalysisCmd.MarkFlagRequired("route-analysis-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getRouteAnalysisCmd)
}
