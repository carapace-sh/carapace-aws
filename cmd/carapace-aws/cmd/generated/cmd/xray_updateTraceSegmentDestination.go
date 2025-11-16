package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_updateTraceSegmentDestinationCmd = &cobra.Command{
	Use:   "update-trace-segment-destination",
	Short: "Modifies the destination of data sent to `PutTraceSegments`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_updateTraceSegmentDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_updateTraceSegmentDestinationCmd).Standalone()

		xray_updateTraceSegmentDestinationCmd.Flags().String("destination", "", "The configured destination of trace segments.")
	})
	xrayCmd.AddCommand(xray_updateTraceSegmentDestinationCmd)
}
