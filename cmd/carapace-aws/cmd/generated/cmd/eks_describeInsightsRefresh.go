package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeInsightsRefreshCmd = &cobra.Command{
	Use:   "describe-insights-refresh",
	Short: "Returns the status of the latest on-demand cluster insights refresh operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeInsightsRefreshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describeInsightsRefreshCmd).Standalone()

		eks_describeInsightsRefreshCmd.Flags().String("cluster-name", "", "The name of the cluster associated with the insights refresh operation.")
		eks_describeInsightsRefreshCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_describeInsightsRefreshCmd)
}
