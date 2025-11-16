package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_getImageSetMetadataCmd = &cobra.Command{
	Use:   "get-image-set-metadata",
	Short: "Get metadata attributes for an image set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_getImageSetMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImaging_getImageSetMetadataCmd).Standalone()

		medicalImaging_getImageSetMetadataCmd.Flags().String("datastore-id", "", "The data store identifier.")
		medicalImaging_getImageSetMetadataCmd.Flags().String("image-set-id", "", "The image set identifier.")
		medicalImaging_getImageSetMetadataCmd.Flags().String("version-id", "", "The image set version identifier.")
		medicalImaging_getImageSetMetadataCmd.MarkFlagRequired("datastore-id")
		medicalImaging_getImageSetMetadataCmd.MarkFlagRequired("image-set-id")
	})
	medicalImagingCmd.AddCommand(medicalImaging_getImageSetMetadataCmd)
}
