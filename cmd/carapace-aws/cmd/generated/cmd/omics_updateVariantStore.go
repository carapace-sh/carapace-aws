package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateVariantStoreCmd = &cobra.Command{
	Use:   "update-variant-store",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateVariantStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_updateVariantStoreCmd).Standalone()

		omics_updateVariantStoreCmd.Flags().String("description", "", "A description for the store.")
		omics_updateVariantStoreCmd.Flags().String("name", "", "A name for the store.")
		omics_updateVariantStoreCmd.MarkFlagRequired("name")
	})
	omicsCmd.AddCommand(omics_updateVariantStoreCmd)
}
