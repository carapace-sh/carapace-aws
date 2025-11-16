package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_listPerformanceAnalysisReportsCmd = &cobra.Command{
	Use:   "list-performance-analysis-reports",
	Short: "Lists all the analysis reports created for the DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_listPerformanceAnalysisReportsCmd).Standalone()

	pi_listPerformanceAnalysisReportsCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique for an Amazon Web Services Region.")
	pi_listPerformanceAnalysisReportsCmd.Flags().Bool("list-tags", false, "Specifies whether or not to include the list of tags in the response.")
	pi_listPerformanceAnalysisReportsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	pi_listPerformanceAnalysisReportsCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous request.")
	pi_listPerformanceAnalysisReportsCmd.Flags().Bool("no-list-tags", false, "Specifies whether or not to include the list of tags in the response.")
	pi_listPerformanceAnalysisReportsCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns metrics.")
	pi_listPerformanceAnalysisReportsCmd.MarkFlagRequired("identifier")
	pi_listPerformanceAnalysisReportsCmd.Flag("no-list-tags").Hidden = true
	pi_listPerformanceAnalysisReportsCmd.MarkFlagRequired("service-type")
	piCmd.AddCommand(pi_listPerformanceAnalysisReportsCmd)
}
