package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteVariantStoreCmd = &cobra.Command{
	Use:   "delete-variant-store",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteVariantStoreCmd).Standalone()

	omics_deleteVariantStoreCmd.Flags().String("force", "", "Whether to force deletion.")
	omics_deleteVariantStoreCmd.Flags().String("name", "", "The store's name.")
	omics_deleteVariantStoreCmd.MarkFlagRequired("name")
	omicsCmd.AddCommand(omics_deleteVariantStoreCmd)
}
