package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_startTraceRetrievalCmd = &cobra.Command{
	Use:   "start-trace-retrieval",
	Short: "Initiates a trace retrieval process using the specified time range and for the given trace IDs in the Transaction Search generated CloudWatch log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_startTraceRetrievalCmd).Standalone()

	xray_startTraceRetrievalCmd.Flags().String("end-time", "", "The end of the time range to retrieve traces.")
	xray_startTraceRetrievalCmd.Flags().String("start-time", "", "The start of the time range to retrieve traces.")
	xray_startTraceRetrievalCmd.Flags().String("trace-ids", "", "Specify the trace IDs of the traces to be retrieved.")
	xray_startTraceRetrievalCmd.MarkFlagRequired("end-time")
	xray_startTraceRetrievalCmd.MarkFlagRequired("start-time")
	xray_startTraceRetrievalCmd.MarkFlagRequired("trace-ids")
	xrayCmd.AddCommand(xray_startTraceRetrievalCmd)
}
