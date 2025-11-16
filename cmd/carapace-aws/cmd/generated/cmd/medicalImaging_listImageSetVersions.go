package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_listImageSetVersionsCmd = &cobra.Command{
	Use:   "list-image-set-versions",
	Short: "List image set versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_listImageSetVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImaging_listImageSetVersionsCmd).Standalone()

		medicalImaging_listImageSetVersionsCmd.Flags().String("datastore-id", "", "The data store identifier.")
		medicalImaging_listImageSetVersionsCmd.Flags().String("image-set-id", "", "The image set identifier.")
		medicalImaging_listImageSetVersionsCmd.Flags().String("max-results", "", "The max results count.")
		medicalImaging_listImageSetVersionsCmd.Flags().String("next-token", "", "The pagination token used to request the list of image set versions on the next page.")
		medicalImaging_listImageSetVersionsCmd.MarkFlagRequired("datastore-id")
		medicalImaging_listImageSetVersionsCmd.MarkFlagRequired("image-set-id")
	})
	medicalImagingCmd.AddCommand(medicalImaging_listImageSetVersionsCmd)
}
