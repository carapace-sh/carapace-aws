package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createAnnotationStoreCmd = &cobra.Command{
	Use:   "create-annotation-store",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createAnnotationStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createAnnotationStoreCmd).Standalone()

		omics_createAnnotationStoreCmd.Flags().String("description", "", "A description for the store.")
		omics_createAnnotationStoreCmd.Flags().String("name", "", "A name for the store.")
		omics_createAnnotationStoreCmd.Flags().String("reference", "", "The genome reference for the store's annotations.")
		omics_createAnnotationStoreCmd.Flags().String("sse-config", "", "Server-side encryption (SSE) settings for the store.")
		omics_createAnnotationStoreCmd.Flags().String("store-format", "", "The annotation file format of the store.")
		omics_createAnnotationStoreCmd.Flags().String("store-options", "", "File parsing options for the annotation store.")
		omics_createAnnotationStoreCmd.Flags().String("tags", "", "Tags for the store.")
		omics_createAnnotationStoreCmd.Flags().String("version-name", "", "The name given to an annotation store version to distinguish it from other versions.")
		omics_createAnnotationStoreCmd.MarkFlagRequired("store-format")
	})
	omicsCmd.AddCommand(omics_createAnnotationStoreCmd)
}
