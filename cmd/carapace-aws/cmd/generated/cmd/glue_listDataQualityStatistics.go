package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listDataQualityStatisticsCmd = &cobra.Command{
	Use:   "list-data-quality-statistics",
	Short: "Retrieves a list of data quality statistics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listDataQualityStatisticsCmd).Standalone()

	glue_listDataQualityStatisticsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	glue_listDataQualityStatisticsCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
	glue_listDataQualityStatisticsCmd.Flags().String("profile-id", "", "The Profile ID.")
	glue_listDataQualityStatisticsCmd.Flags().String("statistic-id", "", "The Statistic ID.")
	glue_listDataQualityStatisticsCmd.Flags().String("timestamp-filter", "", "A timestamp filter.")
	glueCmd.AddCommand(glue_listDataQualityStatisticsCmd)
}
