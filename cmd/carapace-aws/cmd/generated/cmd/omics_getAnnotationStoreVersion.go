package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getAnnotationStoreVersionCmd = &cobra.Command{
	Use:   "get-annotation-store-version",
	Short: "Retrieves the metadata for an annotation store version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getAnnotationStoreVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getAnnotationStoreVersionCmd).Standalone()

		omics_getAnnotationStoreVersionCmd.Flags().String("name", "", "The name given to an annotation store version to distinguish it from others.")
		omics_getAnnotationStoreVersionCmd.Flags().String("version-name", "", "The name given to an annotation store version to distinguish it from others.")
		omics_getAnnotationStoreVersionCmd.MarkFlagRequired("name")
		omics_getAnnotationStoreVersionCmd.MarkFlagRequired("version-name")
	})
	omicsCmd.AddCommand(omics_getAnnotationStoreVersionCmd)
}
