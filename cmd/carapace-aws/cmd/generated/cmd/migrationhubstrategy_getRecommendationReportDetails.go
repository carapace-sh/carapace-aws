package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getRecommendationReportDetailsCmd = &cobra.Command{
	Use:   "get-recommendation-report-details",
	Short: "Retrieves detailed information about the specified recommendation report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getRecommendationReportDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getRecommendationReportDetailsCmd).Standalone()

		migrationhubstrategy_getRecommendationReportDetailsCmd.Flags().String("id", "", "The recommendation report generation task `id` returned by [StartRecommendationReportGeneration]().")
		migrationhubstrategy_getRecommendationReportDetailsCmd.MarkFlagRequired("id")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getRecommendationReportDetailsCmd)
}
