package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_searchFacesCmd = &cobra.Command{
	Use:   "search-faces",
	Short: "For a given input face ID, searches for matching faces in the collection the face belongs to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_searchFacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_searchFacesCmd).Standalone()

		rekognition_searchFacesCmd.Flags().String("collection-id", "", "ID of the collection the face belongs to.")
		rekognition_searchFacesCmd.Flags().String("face-id", "", "ID of a face to find matches for in the collection.")
		rekognition_searchFacesCmd.Flags().String("face-match-threshold", "", "Optional value specifying the minimum confidence in the face match to return.")
		rekognition_searchFacesCmd.Flags().String("max-faces", "", "Maximum number of faces to return.")
		rekognition_searchFacesCmd.MarkFlagRequired("collection-id")
		rekognition_searchFacesCmd.MarkFlagRequired("face-id")
	})
	rekognitionCmd.AddCommand(rekognition_searchFacesCmd)
}
