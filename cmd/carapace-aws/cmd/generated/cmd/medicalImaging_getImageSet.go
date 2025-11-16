package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_getImageSetCmd = &cobra.Command{
	Use:   "get-image-set",
	Short: "Get image set properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_getImageSetCmd).Standalone()

	medicalImaging_getImageSetCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_getImageSetCmd.Flags().String("image-set-id", "", "The image set identifier.")
	medicalImaging_getImageSetCmd.Flags().String("version-id", "", "The image set version identifier.")
	medicalImaging_getImageSetCmd.MarkFlagRequired("datastore-id")
	medicalImaging_getImageSetCmd.MarkFlagRequired("image-set-id")
	medicalImagingCmd.AddCommand(medicalImaging_getImageSetCmd)
}
