package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_batchGetTracesCmd = &cobra.Command{
	Use:   "batch-get-traces",
	Short: "You cannot find traces through this API if Transaction Search is enabled since trace is not indexed in X-Ray.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_batchGetTracesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_batchGetTracesCmd).Standalone()

		xray_batchGetTracesCmd.Flags().String("next-token", "", "Pagination token.")
		xray_batchGetTracesCmd.Flags().String("trace-ids", "", "Specify the trace IDs of requests for which to retrieve segments.")
		xray_batchGetTracesCmd.MarkFlagRequired("trace-ids")
	})
	xrayCmd.AddCommand(xray_batchGetTracesCmd)
}
