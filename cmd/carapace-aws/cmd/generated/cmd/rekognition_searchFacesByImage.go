package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_searchFacesByImageCmd = &cobra.Command{
	Use:   "search-faces-by-image",
	Short: "For a given input image, first detects the largest face in the image, and then searches the specified collection for matching faces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_searchFacesByImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_searchFacesByImageCmd).Standalone()

		rekognition_searchFacesByImageCmd.Flags().String("collection-id", "", "ID of the collection to search.")
		rekognition_searchFacesByImageCmd.Flags().String("face-match-threshold", "", "(Optional) Specifies the minimum confidence in the face match to return.")
		rekognition_searchFacesByImageCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an S3 object.")
		rekognition_searchFacesByImageCmd.Flags().String("max-faces", "", "Maximum number of faces to return.")
		rekognition_searchFacesByImageCmd.Flags().String("quality-filter", "", "A filter that specifies a quality bar for how much filtering is done to identify faces.")
		rekognition_searchFacesByImageCmd.MarkFlagRequired("collection-id")
		rekognition_searchFacesByImageCmd.MarkFlagRequired("image")
	})
	rekognitionCmd.AddCommand(rekognition_searchFacesByImageCmd)
}
