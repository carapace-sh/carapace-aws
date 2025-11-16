package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getConsolidatedReportCmd = &cobra.Command{
	Use:   "get-consolidated-report",
	Short: "Get a consolidated report of your workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getConsolidatedReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getConsolidatedReportCmd).Standalone()

		wellarchitected_getConsolidatedReportCmd.Flags().String("format", "", "The format of the consolidated report.")
		wellarchitected_getConsolidatedReportCmd.Flags().String("include-shared-resources", "", "Set to `true` to have shared resources included in the report.")
		wellarchitected_getConsolidatedReportCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
		wellarchitected_getConsolidatedReportCmd.Flags().String("next-token", "", "")
		wellarchitected_getConsolidatedReportCmd.MarkFlagRequired("format")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_getConsolidatedReportCmd)
}
