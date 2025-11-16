package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listDataQualityResultsCmd = &cobra.Command{
	Use:   "list-data-quality-results",
	Short: "Returns all data quality execution results for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listDataQualityResultsCmd).Standalone()

	glue_listDataQualityResultsCmd.Flags().String("filter", "", "The filter criteria.")
	glue_listDataQualityResultsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	glue_listDataQualityResultsCmd.Flags().String("next-token", "", "A paginated token to offset the results.")
	glueCmd.AddCommand(glue_listDataQualityResultsCmd)
}
