package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReadSetExportJobCmd = &cobra.Command{
	Use:   "get-read-set-export-job",
	Short: "Retrieves status information about a read set export job and returns the data in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReadSetExportJobCmd).Standalone()

	omics_getReadSetExportJobCmd.Flags().String("id", "", "The job's ID.")
	omics_getReadSetExportJobCmd.Flags().String("sequence-store-id", "", "The job's sequence store ID.")
	omics_getReadSetExportJobCmd.MarkFlagRequired("id")
	omics_getReadSetExportJobCmd.MarkFlagRequired("sequence-store-id")
	omicsCmd.AddCommand(omics_getReadSetExportJobCmd)
}
