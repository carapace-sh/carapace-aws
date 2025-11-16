package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getSamplingStatisticSummariesCmd = &cobra.Command{
	Use:   "get-sampling-statistic-summaries",
	Short: "Retrieves information about recent sampling results for all sampling rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getSamplingStatisticSummariesCmd).Standalone()

	xray_getSamplingStatisticSummariesCmd.Flags().String("next-token", "", "Pagination token.")
	xrayCmd.AddCommand(xray_getSamplingStatisticSummariesCmd)
}
