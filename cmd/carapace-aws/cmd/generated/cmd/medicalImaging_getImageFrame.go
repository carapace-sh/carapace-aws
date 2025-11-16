package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_getImageFrameCmd = &cobra.Command{
	Use:   "get-image-frame",
	Short: "Get an image frame (pixel data) for an image set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_getImageFrameCmd).Standalone()

	medicalImaging_getImageFrameCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_getImageFrameCmd.Flags().String("image-frame-information", "", "Information about the image frame (pixel data) identifier.")
	medicalImaging_getImageFrameCmd.Flags().String("image-set-id", "", "The image set identifier.")
	medicalImaging_getImageFrameCmd.MarkFlagRequired("datastore-id")
	medicalImaging_getImageFrameCmd.MarkFlagRequired("image-frame-information")
	medicalImaging_getImageFrameCmd.MarkFlagRequired("image-set-id")
	medicalImagingCmd.AddCommand(medicalImaging_getImageFrameCmd)
}
