package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listDataQualityStatisticAnnotationsCmd = &cobra.Command{
	Use:   "list-data-quality-statistic-annotations",
	Short: "Retrieve annotations for a data quality statistic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listDataQualityStatisticAnnotationsCmd).Standalone()

	glue_listDataQualityStatisticAnnotationsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	glue_listDataQualityStatisticAnnotationsCmd.Flags().String("next-token", "", "A pagination token to retrieve the next set of results.")
	glue_listDataQualityStatisticAnnotationsCmd.Flags().String("profile-id", "", "The Profile ID.")
	glue_listDataQualityStatisticAnnotationsCmd.Flags().String("statistic-id", "", "The Statistic ID.")
	glue_listDataQualityStatisticAnnotationsCmd.Flags().String("timestamp-filter", "", "A timestamp filter.")
	glueCmd.AddCommand(glue_listDataQualityStatisticAnnotationsCmd)
}
