package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReferenceImportJobCmd = &cobra.Command{
	Use:   "get-reference-import-job",
	Short: "Monitors the status of a reference import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReferenceImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getReferenceImportJobCmd).Standalone()

		omics_getReferenceImportJobCmd.Flags().String("id", "", "The job's ID.")
		omics_getReferenceImportJobCmd.Flags().String("reference-store-id", "", "The job's reference store ID.")
		omics_getReferenceImportJobCmd.MarkFlagRequired("id")
		omics_getReferenceImportJobCmd.MarkFlagRequired("reference-store-id")
	})
	omicsCmd.AddCommand(omics_getReferenceImportJobCmd)
}
