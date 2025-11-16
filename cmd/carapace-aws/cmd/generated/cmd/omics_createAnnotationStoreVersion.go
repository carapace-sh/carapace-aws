package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createAnnotationStoreVersionCmd = &cobra.Command{
	Use:   "create-annotation-store-version",
	Short: "Creates a new version of an annotation store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createAnnotationStoreVersionCmd).Standalone()

	omics_createAnnotationStoreVersionCmd.Flags().String("description", "", "The description of an annotation store version.")
	omics_createAnnotationStoreVersionCmd.Flags().String("name", "", "The name of an annotation store version from which versions are being created.")
	omics_createAnnotationStoreVersionCmd.Flags().String("tags", "", "Any tags added to annotation store version.")
	omics_createAnnotationStoreVersionCmd.Flags().String("version-name", "", "The name given to an annotation store version to distinguish it from other versions.")
	omics_createAnnotationStoreVersionCmd.Flags().String("version-options", "", "The options for an annotation store version.")
	omics_createAnnotationStoreVersionCmd.MarkFlagRequired("name")
	omics_createAnnotationStoreVersionCmd.MarkFlagRequired("version-name")
	omicsCmd.AddCommand(omics_createAnnotationStoreVersionCmd)
}
