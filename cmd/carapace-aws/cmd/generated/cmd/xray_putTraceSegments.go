package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_putTraceSegmentsCmd = &cobra.Command{
	Use:   "put-trace-segments",
	Short: "Uploads segment documents to Amazon Web Services X-Ray.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_putTraceSegmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_putTraceSegmentsCmd).Standalone()

		xray_putTraceSegmentsCmd.Flags().String("trace-segment-documents", "", "A string containing a JSON document defining one or more segments or subsegments.")
		xray_putTraceSegmentsCmd.MarkFlagRequired("trace-segment-documents")
	})
	xrayCmd.AddCommand(xray_putTraceSegmentsCmd)
}
