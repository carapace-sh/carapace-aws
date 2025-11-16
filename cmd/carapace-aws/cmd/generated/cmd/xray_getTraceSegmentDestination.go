package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getTraceSegmentDestinationCmd = &cobra.Command{
	Use:   "get-trace-segment-destination",
	Short: "Retrieves the current destination of data sent to `PutTraceSegments` and *OpenTelemetry protocol (OTLP)* endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getTraceSegmentDestinationCmd).Standalone()

	xrayCmd.AddCommand(xray_getTraceSegmentDestinationCmd)
}
