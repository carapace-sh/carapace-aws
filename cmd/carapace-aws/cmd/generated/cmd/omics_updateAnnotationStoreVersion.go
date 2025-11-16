package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateAnnotationStoreVersionCmd = &cobra.Command{
	Use:   "update-annotation-store-version",
	Short: "Updates the description of an annotation store version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateAnnotationStoreVersionCmd).Standalone()

	omics_updateAnnotationStoreVersionCmd.Flags().String("description", "", "The description of an annotation store.")
	omics_updateAnnotationStoreVersionCmd.Flags().String("name", "", "The name of an annotation store.")
	omics_updateAnnotationStoreVersionCmd.Flags().String("version-name", "", "The name of an annotation store version.")
	omics_updateAnnotationStoreVersionCmd.MarkFlagRequired("name")
	omics_updateAnnotationStoreVersionCmd.MarkFlagRequired("version-name")
	omicsCmd.AddCommand(omics_updateAnnotationStoreVersionCmd)
}
