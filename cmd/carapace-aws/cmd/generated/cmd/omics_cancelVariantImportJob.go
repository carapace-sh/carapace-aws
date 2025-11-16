package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_cancelVariantImportJobCmd = &cobra.Command{
	Use:   "cancel-variant-import-job",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_cancelVariantImportJobCmd).Standalone()

	omics_cancelVariantImportJobCmd.Flags().String("job-id", "", "The job's ID.")
	omics_cancelVariantImportJobCmd.MarkFlagRequired("job-id")
	omicsCmd.AddCommand(omics_cancelVariantImportJobCmd)
}
