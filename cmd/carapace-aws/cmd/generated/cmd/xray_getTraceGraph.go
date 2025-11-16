package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getTraceGraphCmd = &cobra.Command{
	Use:   "get-trace-graph",
	Short: "Retrieves a service graph for one or more specific trace IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getTraceGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_getTraceGraphCmd).Standalone()

		xray_getTraceGraphCmd.Flags().String("next-token", "", "Pagination token.")
		xray_getTraceGraphCmd.Flags().String("trace-ids", "", "Trace IDs of requests for which to generate a service graph.")
		xray_getTraceGraphCmd.MarkFlagRequired("trace-ids")
	})
	xrayCmd.AddCommand(xray_getTraceGraphCmd)
}
