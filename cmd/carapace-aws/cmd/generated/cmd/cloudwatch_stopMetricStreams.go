package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_stopMetricStreamsCmd = &cobra.Command{
	Use:   "stop-metric-streams",
	Short: "Stops the streaming of metrics for one or more of your metric streams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_stopMetricStreamsCmd).Standalone()

	cloudwatch_stopMetricStreamsCmd.Flags().String("names", "", "The array of the names of metric streams to stop streaming.")
	cloudwatch_stopMetricStreamsCmd.MarkFlagRequired("names")
	cloudwatchCmd.AddCommand(cloudwatch_stopMetricStreamsCmd)
}
