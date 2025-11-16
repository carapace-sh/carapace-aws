package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofiler_deleteReportDefinitionCmd = &cobra.Command{
	Use:   "delete-report-definition",
	Short: "Deletes the specified report definition in AWS Application Cost Profiler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofiler_deleteReportDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationcostprofiler_deleteReportDefinitionCmd).Standalone()

		applicationcostprofiler_deleteReportDefinitionCmd.Flags().String("report-id", "", "Required.")
		applicationcostprofiler_deleteReportDefinitionCmd.MarkFlagRequired("report-id")
	})
	applicationcostprofilerCmd.AddCommand(applicationcostprofiler_deleteReportDefinitionCmd)
}
