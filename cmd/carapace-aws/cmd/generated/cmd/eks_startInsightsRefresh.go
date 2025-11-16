package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_startInsightsRefreshCmd = &cobra.Command{
	Use:   "start-insights-refresh",
	Short: "Initiates an on-demand refresh operation for cluster insights, getting the latest analysis outside of the standard refresh schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_startInsightsRefreshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_startInsightsRefreshCmd).Standalone()

		eks_startInsightsRefreshCmd.Flags().String("cluster-name", "", "The name of the cluster for the refresh insights operation.")
		eks_startInsightsRefreshCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_startInsightsRefreshCmd)
}
