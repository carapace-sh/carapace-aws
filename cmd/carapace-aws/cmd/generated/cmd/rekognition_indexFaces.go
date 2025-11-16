package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_indexFacesCmd = &cobra.Command{
	Use:   "index-faces",
	Short: "Detects faces in the input image and adds them to the specified collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_indexFacesCmd).Standalone()

	rekognition_indexFacesCmd.Flags().String("collection-id", "", "The ID of an existing collection to which you want to add the faces that are detected in the input images.")
	rekognition_indexFacesCmd.Flags().String("detection-attributes", "", "An array of facial attributes you want to be returned.")
	rekognition_indexFacesCmd.Flags().String("external-image-id", "", "The ID you want to assign to all the faces detected in the image.")
	rekognition_indexFacesCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an S3 object.")
	rekognition_indexFacesCmd.Flags().String("max-faces", "", "The maximum number of faces to index.")
	rekognition_indexFacesCmd.Flags().String("quality-filter", "", "A filter that specifies a quality bar for how much filtering is done to identify faces.")
	rekognition_indexFacesCmd.MarkFlagRequired("collection-id")
	rekognition_indexFacesCmd.MarkFlagRequired("image")
	rekognitionCmd.AddCommand(rekognition_indexFacesCmd)
}
