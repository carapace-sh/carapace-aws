package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateAnnotationStoreCmd = &cobra.Command{
	Use:   "update-annotation-store",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateAnnotationStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_updateAnnotationStoreCmd).Standalone()

		omics_updateAnnotationStoreCmd.Flags().String("description", "", "A description for the store.")
		omics_updateAnnotationStoreCmd.Flags().String("name", "", "A name for the store.")
		omics_updateAnnotationStoreCmd.MarkFlagRequired("name")
	})
	omicsCmd.AddCommand(omics_updateAnnotationStoreCmd)
}
