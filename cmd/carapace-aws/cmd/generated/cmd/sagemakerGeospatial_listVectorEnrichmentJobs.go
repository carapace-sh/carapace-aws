package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_listVectorEnrichmentJobsCmd = &cobra.Command{
	Use:   "list-vector-enrichment-jobs",
	Short: "Retrieves a list of vector enrichment jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_listVectorEnrichmentJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_listVectorEnrichmentJobsCmd).Standalone()

		sagemakerGeospatial_listVectorEnrichmentJobsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		sagemakerGeospatial_listVectorEnrichmentJobsCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
		sagemakerGeospatial_listVectorEnrichmentJobsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
		sagemakerGeospatial_listVectorEnrichmentJobsCmd.Flags().String("sort-order", "", "An optional value that specifies whether you want the results sorted in `Ascending` or `Descending` order.")
		sagemakerGeospatial_listVectorEnrichmentJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only jobs with a specific status.")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_listVectorEnrichmentJobsCmd)
}
