package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReadSetImportJobCmd = &cobra.Command{
	Use:   "get-read-set-import-job",
	Short: "Gets detailed and status information about a read set import job and returns the data in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReadSetImportJobCmd).Standalone()

	omics_getReadSetImportJobCmd.Flags().String("id", "", "The job's ID.")
	omics_getReadSetImportJobCmd.Flags().String("sequence-store-id", "", "The job's sequence store ID.")
	omics_getReadSetImportJobCmd.MarkFlagRequired("id")
	omics_getReadSetImportJobCmd.MarkFlagRequired("sequence-store-id")
	omicsCmd.AddCommand(omics_getReadSetImportJobCmd)
}
