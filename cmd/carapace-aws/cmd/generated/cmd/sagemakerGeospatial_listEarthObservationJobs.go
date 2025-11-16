package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_listEarthObservationJobsCmd = &cobra.Command{
	Use:   "list-earth-observation-jobs",
	Short: "Use this operation to get a list of the Earth Observation jobs associated with the calling Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_listEarthObservationJobsCmd).Standalone()

	sagemakerGeospatial_listEarthObservationJobsCmd.Flags().String("max-results", "", "The total number of items to return.")
	sagemakerGeospatial_listEarthObservationJobsCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
	sagemakerGeospatial_listEarthObservationJobsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
	sagemakerGeospatial_listEarthObservationJobsCmd.Flags().String("sort-order", "", "An optional value that specifies whether you want the results sorted in `Ascending` or `Descending` order.")
	sagemakerGeospatial_listEarthObservationJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only jobs with a specific status.")
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_listEarthObservationJobsCmd)
}
