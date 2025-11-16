package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofiler_updateReportDefinitionCmd = &cobra.Command{
	Use:   "update-report-definition",
	Short: "Updates existing report in AWS Application Cost Profiler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofiler_updateReportDefinitionCmd).Standalone()

	applicationcostprofiler_updateReportDefinitionCmd.Flags().String("destination-s3-location", "", "Required.")
	applicationcostprofiler_updateReportDefinitionCmd.Flags().String("format", "", "Required.")
	applicationcostprofiler_updateReportDefinitionCmd.Flags().String("report-description", "", "Required.")
	applicationcostprofiler_updateReportDefinitionCmd.Flags().String("report-frequency", "", "Required.")
	applicationcostprofiler_updateReportDefinitionCmd.Flags().String("report-id", "", "Required.")
	applicationcostprofiler_updateReportDefinitionCmd.MarkFlagRequired("destination-s3-location")
	applicationcostprofiler_updateReportDefinitionCmd.MarkFlagRequired("format")
	applicationcostprofiler_updateReportDefinitionCmd.MarkFlagRequired("report-description")
	applicationcostprofiler_updateReportDefinitionCmd.MarkFlagRequired("report-frequency")
	applicationcostprofiler_updateReportDefinitionCmd.MarkFlagRequired("report-id")
	applicationcostprofilerCmd.AddCommand(applicationcostprofiler_updateReportDefinitionCmd)
}
