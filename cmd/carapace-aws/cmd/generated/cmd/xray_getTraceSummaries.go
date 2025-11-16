package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getTraceSummariesCmd = &cobra.Command{
	Use:   "get-trace-summaries",
	Short: "Retrieves IDs and annotations for traces available for a specified time frame using an optional filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getTraceSummariesCmd).Standalone()

	xray_getTraceSummariesCmd.Flags().String("end-time", "", "The end of the time frame for which to retrieve traces.")
	xray_getTraceSummariesCmd.Flags().String("filter-expression", "", "Specify a filter expression to retrieve trace summaries for services or requests that meet certain requirements.")
	xray_getTraceSummariesCmd.Flags().String("next-token", "", "Specify the pagination token returned by a previous request to retrieve the next page of results.")
	xray_getTraceSummariesCmd.Flags().String("sampling", "", "Set to `true` to get summaries for only a subset of available traces.")
	xray_getTraceSummariesCmd.Flags().String("sampling-strategy", "", "A parameter to indicate whether to enable sampling on trace summaries.")
	xray_getTraceSummariesCmd.Flags().String("start-time", "", "The start of the time frame for which to retrieve traces.")
	xray_getTraceSummariesCmd.Flags().String("time-range-type", "", "Query trace summaries by TraceId (trace start time), Event (trace update time), or Service (trace segment end time).")
	xray_getTraceSummariesCmd.MarkFlagRequired("end-time")
	xray_getTraceSummariesCmd.MarkFlagRequired("start-time")
	xrayCmd.AddCommand(xray_getTraceSummariesCmd)
}
