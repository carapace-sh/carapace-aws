package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_detectFacesCmd = &cobra.Command{
	Use:   "detect-faces",
	Short: "Detects faces within an image that is provided as input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_detectFacesCmd).Standalone()

	rekognition_detectFacesCmd.Flags().String("attributes", "", "An array of facial attributes you want to be returned.")
	rekognition_detectFacesCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an S3 object.")
	rekognition_detectFacesCmd.MarkFlagRequired("image")
	rekognitionCmd.AddCommand(rekognition_detectFacesCmd)
}
