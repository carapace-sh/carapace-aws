package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_getMetricStreamCmd = &cobra.Command{
	Use:   "get-metric-stream",
	Short: "Returns information about the metric stream that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_getMetricStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_getMetricStreamCmd).Standalone()

		cloudwatch_getMetricStreamCmd.Flags().String("name", "", "The name of the metric stream to retrieve information about.")
		cloudwatch_getMetricStreamCmd.MarkFlagRequired("name")
	})
	cloudwatchCmd.AddCommand(cloudwatch_getMetricStreamCmd)
}
