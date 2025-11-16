package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_startRecommendationReportGenerationCmd = &cobra.Command{
	Use:   "start-recommendation-report-generation",
	Short: "Starts generating a recommendation report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_startRecommendationReportGenerationCmd).Standalone()

	migrationhubstrategy_startRecommendationReportGenerationCmd.Flags().String("group-id-filter", "", "Groups the resources in the recommendation report with a unique name.")
	migrationhubstrategy_startRecommendationReportGenerationCmd.Flags().String("output-format", "", "The output format for the recommendation report file.")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_startRecommendationReportGenerationCmd)
}
