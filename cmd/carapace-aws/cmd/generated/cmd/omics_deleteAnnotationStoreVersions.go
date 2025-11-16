package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteAnnotationStoreVersionsCmd = &cobra.Command{
	Use:   "delete-annotation-store-versions",
	Short: "Deletes one or multiple versions of an annotation store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteAnnotationStoreVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteAnnotationStoreVersionsCmd).Standalone()

		omics_deleteAnnotationStoreVersionsCmd.Flags().String("force", "", "Forces the deletion of an annotation store version when imports are in-progress..")
		omics_deleteAnnotationStoreVersionsCmd.Flags().String("name", "", "The name of the annotation store from which versions are being deleted.")
		omics_deleteAnnotationStoreVersionsCmd.Flags().String("versions", "", "The versions of an annotation store to be deleted.")
		omics_deleteAnnotationStoreVersionsCmd.MarkFlagRequired("name")
		omics_deleteAnnotationStoreVersionsCmd.MarkFlagRequired("versions")
	})
	omicsCmd.AddCommand(omics_deleteAnnotationStoreVersionsCmd)
}
