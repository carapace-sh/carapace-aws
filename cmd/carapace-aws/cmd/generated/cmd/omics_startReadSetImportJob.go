package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_startReadSetImportJobCmd = &cobra.Command{
	Use:   "start-read-set-import-job",
	Short: "Imports a read set from the sequence store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_startReadSetImportJobCmd).Standalone()

	omics_startReadSetImportJobCmd.Flags().String("client-token", "", "To ensure that jobs don't run multiple times, specify a unique token for each job.")
	omics_startReadSetImportJobCmd.Flags().String("role-arn", "", "A service role for the job.")
	omics_startReadSetImportJobCmd.Flags().String("sequence-store-id", "", "The read set's sequence store ID.")
	omics_startReadSetImportJobCmd.Flags().String("sources", "", "The job's source files.")
	omics_startReadSetImportJobCmd.MarkFlagRequired("role-arn")
	omics_startReadSetImportJobCmd.MarkFlagRequired("sequence-store-id")
	omics_startReadSetImportJobCmd.MarkFlagRequired("sources")
	omicsCmd.AddCommand(omics_startReadSetImportJobCmd)
}
