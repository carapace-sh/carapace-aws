package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_deleteImageSetCmd = &cobra.Command{
	Use:   "delete-image-set",
	Short: "Delete an image set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_deleteImageSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImaging_deleteImageSetCmd).Standalone()

		medicalImaging_deleteImageSetCmd.Flags().String("datastore-id", "", "The data store identifier.")
		medicalImaging_deleteImageSetCmd.Flags().String("image-set-id", "", "The image set identifier.")
		medicalImaging_deleteImageSetCmd.MarkFlagRequired("datastore-id")
		medicalImaging_deleteImageSetCmd.MarkFlagRequired("image-set-id")
	})
	medicalImagingCmd.AddCommand(medicalImaging_deleteImageSetCmd)
}
