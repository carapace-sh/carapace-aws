package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_updateImageSetMetadataCmd = &cobra.Command{
	Use:   "update-image-set-metadata",
	Short: "Update image set metadata attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_updateImageSetMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImaging_updateImageSetMetadataCmd).Standalone()

		medicalImaging_updateImageSetMetadataCmd.Flags().String("datastore-id", "", "The data store identifier.")
		medicalImaging_updateImageSetMetadataCmd.Flags().Bool("force", false, "Setting this flag will force the `UpdateImageSetMetadata` operation for the following attributes:")
		medicalImaging_updateImageSetMetadataCmd.Flags().String("image-set-id", "", "The image set identifier.")
		medicalImaging_updateImageSetMetadataCmd.Flags().String("latest-version-id", "", "The latest image set version identifier.")
		medicalImaging_updateImageSetMetadataCmd.Flags().Bool("no-force", false, "Setting this flag will force the `UpdateImageSetMetadata` operation for the following attributes:")
		medicalImaging_updateImageSetMetadataCmd.Flags().String("update-image-set-metadata-updates", "", "Update image set metadata updates.")
		medicalImaging_updateImageSetMetadataCmd.MarkFlagRequired("datastore-id")
		medicalImaging_updateImageSetMetadataCmd.MarkFlagRequired("image-set-id")
		medicalImaging_updateImageSetMetadataCmd.MarkFlagRequired("latest-version-id")
		medicalImaging_updateImageSetMetadataCmd.Flag("no-force").Hidden = true
		medicalImaging_updateImageSetMetadataCmd.MarkFlagRequired("update-image-set-metadata-updates")
	})
	medicalImagingCmd.AddCommand(medicalImaging_updateImageSetMetadataCmd)
}
