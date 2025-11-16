package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofiler_listReportDefinitionsCmd = &cobra.Command{
	Use:   "list-report-definitions",
	Short: "Retrieves a list of all reports and their configurations for your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofiler_listReportDefinitionsCmd).Standalone()

	applicationcostprofiler_listReportDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	applicationcostprofiler_listReportDefinitionsCmd.Flags().String("next-token", "", "The token value from a previous call to access the next page of results.")
	applicationcostprofilerCmd.AddCommand(applicationcostprofiler_listReportDefinitionsCmd)
}
