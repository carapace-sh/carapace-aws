package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_listMetricStreamsCmd = &cobra.Command{
	Use:   "list-metric-streams",
	Short: "Returns a list of metric streams in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_listMetricStreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_listMetricStreamsCmd).Standalone()

		cloudwatch_listMetricStreamsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
		cloudwatch_listMetricStreamsCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous call, to get the next set of metric streams.")
	})
	cloudwatchCmd.AddCommand(cloudwatch_listMetricStreamsCmd)
}
