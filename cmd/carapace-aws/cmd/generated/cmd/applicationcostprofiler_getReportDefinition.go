package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofiler_getReportDefinitionCmd = &cobra.Command{
	Use:   "get-report-definition",
	Short: "Retrieves the definition of a report already configured in AWS Application Cost Profiler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofiler_getReportDefinitionCmd).Standalone()

	applicationcostprofiler_getReportDefinitionCmd.Flags().String("report-id", "", "ID of the report to retrieve.")
	applicationcostprofiler_getReportDefinitionCmd.MarkFlagRequired("report-id")
	applicationcostprofilerCmd.AddCommand(applicationcostprofiler_getReportDefinitionCmd)
}
