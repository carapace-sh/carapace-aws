package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_searchImageSetsCmd = &cobra.Command{
	Use:   "search-image-sets",
	Short: "Search image sets based on defined input attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_searchImageSetsCmd).Standalone()

	medicalImaging_searchImageSetsCmd.Flags().String("datastore-id", "", "The identifier of the data store where the image sets reside.")
	medicalImaging_searchImageSetsCmd.Flags().String("max-results", "", "The maximum number of results that can be returned in a search.")
	medicalImaging_searchImageSetsCmd.Flags().String("next-token", "", "The token used for pagination of results returned in the response.")
	medicalImaging_searchImageSetsCmd.Flags().String("search-criteria", "", "The search criteria that filters by applying a maximum of 1 item to `SearchByAttribute`.")
	medicalImaging_searchImageSetsCmd.MarkFlagRequired("datastore-id")
	medicalImagingCmd.AddCommand(medicalImaging_searchImageSetsCmd)
}
