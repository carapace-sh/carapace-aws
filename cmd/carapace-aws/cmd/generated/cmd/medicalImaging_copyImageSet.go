package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_copyImageSetCmd = &cobra.Command{
	Use:   "copy-image-set",
	Short: "Copy an image set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_copyImageSetCmd).Standalone()

	medicalImaging_copyImageSetCmd.Flags().String("copy-image-set-information", "", "Copy image set information.")
	medicalImaging_copyImageSetCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_copyImageSetCmd.Flags().Bool("force", false, "Providing this parameter will force completion of the `CopyImageSet` operation, even if there are inconsistent Patient, Study, and/or Series level metadata elements between the `sourceImageSet` and `destinationImageSet`.")
	medicalImaging_copyImageSetCmd.Flags().Bool("no-force", false, "Providing this parameter will force completion of the `CopyImageSet` operation, even if there are inconsistent Patient, Study, and/or Series level metadata elements between the `sourceImageSet` and `destinationImageSet`.")
	medicalImaging_copyImageSetCmd.Flags().Bool("no-promote-to-primary", false, "Providing this parameter will configure the `CopyImageSet` operation to promote the given image set to the primary DICOM hierarchy.")
	medicalImaging_copyImageSetCmd.Flags().Bool("promote-to-primary", false, "Providing this parameter will configure the `CopyImageSet` operation to promote the given image set to the primary DICOM hierarchy.")
	medicalImaging_copyImageSetCmd.Flags().String("source-image-set-id", "", "The source image set identifier.")
	medicalImaging_copyImageSetCmd.MarkFlagRequired("copy-image-set-information")
	medicalImaging_copyImageSetCmd.MarkFlagRequired("datastore-id")
	medicalImaging_copyImageSetCmd.Flag("no-force").Hidden = true
	medicalImaging_copyImageSetCmd.Flag("no-promote-to-primary").Hidden = true
	medicalImaging_copyImageSetCmd.MarkFlagRequired("source-image-set-id")
	medicalImagingCmd.AddCommand(medicalImaging_copyImageSetCmd)
}
