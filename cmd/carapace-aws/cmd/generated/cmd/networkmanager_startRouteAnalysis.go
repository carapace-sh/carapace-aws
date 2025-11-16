package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_startRouteAnalysisCmd = &cobra.Command{
	Use:   "start-route-analysis",
	Short: "Starts analyzing the routing path between the specified source and destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_startRouteAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_startRouteAnalysisCmd).Standalone()

		networkmanager_startRouteAnalysisCmd.Flags().String("destination", "", "The destination.")
		networkmanager_startRouteAnalysisCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_startRouteAnalysisCmd.Flags().Bool("include-return-path", false, "Indicates whether to analyze the return path.")
		networkmanager_startRouteAnalysisCmd.Flags().Bool("no-include-return-path", false, "Indicates whether to analyze the return path.")
		networkmanager_startRouteAnalysisCmd.Flags().Bool("no-use-middleboxes", false, "Indicates whether to include the location of middlebox appliances in the route analysis.")
		networkmanager_startRouteAnalysisCmd.Flags().String("source", "", "The source from which traffic originates.")
		networkmanager_startRouteAnalysisCmd.Flags().Bool("use-middleboxes", false, "Indicates whether to include the location of middlebox appliances in the route analysis.")
		networkmanager_startRouteAnalysisCmd.MarkFlagRequired("destination")
		networkmanager_startRouteAnalysisCmd.MarkFlagRequired("global-network-id")
		networkmanager_startRouteAnalysisCmd.Flag("no-include-return-path").Hidden = true
		networkmanager_startRouteAnalysisCmd.Flag("no-use-middleboxes").Hidden = true
		networkmanager_startRouteAnalysisCmd.MarkFlagRequired("source")
	})
	networkmanagerCmd.AddCommand(networkmanager_startRouteAnalysisCmd)
}
