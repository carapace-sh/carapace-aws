package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_startMetricStreamsCmd = &cobra.Command{
	Use:   "start-metric-streams",
	Short: "Starts the streaming of metrics for one or more of your metric streams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_startMetricStreamsCmd).Standalone()

	cloudwatch_startMetricStreamsCmd.Flags().String("names", "", "The array of the names of metric streams to start streaming.")
	cloudwatch_startMetricStreamsCmd.MarkFlagRequired("names")
	cloudwatchCmd.AddCommand(cloudwatch_startMetricStreamsCmd)
}
