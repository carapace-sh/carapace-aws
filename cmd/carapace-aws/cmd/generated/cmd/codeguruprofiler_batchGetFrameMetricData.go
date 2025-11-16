package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_batchGetFrameMetricDataCmd = &cobra.Command{
	Use:   "batch-get-frame-metric-data",
	Short: "Returns the time series of values for a requested list of frame metrics from a time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_batchGetFrameMetricDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_batchGetFrameMetricDataCmd).Standalone()

		codeguruprofiler_batchGetFrameMetricDataCmd.Flags().String("end-time", "", "The end time of the time period for the returned time series values.")
		codeguruprofiler_batchGetFrameMetricDataCmd.Flags().String("frame-metrics", "", "The details of the metrics that are used to request a time series of values.")
		codeguruprofiler_batchGetFrameMetricDataCmd.Flags().String("period", "", "The duration of the frame metrics used to return the time series values.")
		codeguruprofiler_batchGetFrameMetricDataCmd.Flags().String("profiling-group-name", "", "The name of the profiling group associated with the the frame metrics used to return the time series values.")
		codeguruprofiler_batchGetFrameMetricDataCmd.Flags().String("start-time", "", "The start time of the time period for the frame metrics used to return the time series values.")
		codeguruprofiler_batchGetFrameMetricDataCmd.Flags().String("target-resolution", "", "The requested resolution of time steps for the returned time series of values.")
		codeguruprofiler_batchGetFrameMetricDataCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_batchGetFrameMetricDataCmd)
}
