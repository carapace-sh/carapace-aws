package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_deleteMetricStreamCmd = &cobra.Command{
	Use:   "delete-metric-stream",
	Short: "Permanently deletes the metric stream that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_deleteMetricStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_deleteMetricStreamCmd).Standalone()

		cloudwatch_deleteMetricStreamCmd.Flags().String("name", "", "The name of the metric stream to delete.")
		cloudwatch_deleteMetricStreamCmd.MarkFlagRequired("name")
	})
	cloudwatchCmd.AddCommand(cloudwatch_deleteMetricStreamCmd)
}
