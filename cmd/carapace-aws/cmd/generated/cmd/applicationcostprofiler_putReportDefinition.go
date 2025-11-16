package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofiler_putReportDefinitionCmd = &cobra.Command{
	Use:   "put-report-definition",
	Short: "Creates the report definition for a report in Application Cost Profiler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofiler_putReportDefinitionCmd).Standalone()

	applicationcostprofiler_putReportDefinitionCmd.Flags().String("destination-s3-location", "", "Required.")
	applicationcostprofiler_putReportDefinitionCmd.Flags().String("format", "", "Required.")
	applicationcostprofiler_putReportDefinitionCmd.Flags().String("report-description", "", "Required.")
	applicationcostprofiler_putReportDefinitionCmd.Flags().String("report-frequency", "", "Required.")
	applicationcostprofiler_putReportDefinitionCmd.Flags().String("report-id", "", "Required.")
	applicationcostprofiler_putReportDefinitionCmd.MarkFlagRequired("destination-s3-location")
	applicationcostprofiler_putReportDefinitionCmd.MarkFlagRequired("format")
	applicationcostprofiler_putReportDefinitionCmd.MarkFlagRequired("report-description")
	applicationcostprofiler_putReportDefinitionCmd.MarkFlagRequired("report-frequency")
	applicationcostprofiler_putReportDefinitionCmd.MarkFlagRequired("report-id")
	applicationcostprofilerCmd.AddCommand(applicationcostprofiler_putReportDefinitionCmd)
}
