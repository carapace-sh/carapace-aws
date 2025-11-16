package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createVariantStoreCmd = &cobra.Command{
	Use:   "create-variant-store",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createVariantStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createVariantStoreCmd).Standalone()

		omics_createVariantStoreCmd.Flags().String("description", "", "A description for the store.")
		omics_createVariantStoreCmd.Flags().String("name", "", "A name for the store.")
		omics_createVariantStoreCmd.Flags().String("reference", "", "The genome reference for the store's variants.")
		omics_createVariantStoreCmd.Flags().String("sse-config", "", "Server-side encryption (SSE) settings for the store.")
		omics_createVariantStoreCmd.Flags().String("tags", "", "Tags for the store.")
		omics_createVariantStoreCmd.MarkFlagRequired("reference")
	})
	omicsCmd.AddCommand(omics_createVariantStoreCmd)
}
