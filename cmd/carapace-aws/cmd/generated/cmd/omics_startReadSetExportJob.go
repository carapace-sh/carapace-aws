package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_startReadSetExportJobCmd = &cobra.Command{
	Use:   "start-read-set-export-job",
	Short: "Starts a read set export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_startReadSetExportJobCmd).Standalone()

	omics_startReadSetExportJobCmd.Flags().String("client-token", "", "To ensure that jobs don't run multiple times, specify a unique token for each job.")
	omics_startReadSetExportJobCmd.Flags().String("destination", "", "A location for exported files in Amazon S3.")
	omics_startReadSetExportJobCmd.Flags().String("role-arn", "", "A service role for the job.")
	omics_startReadSetExportJobCmd.Flags().String("sequence-store-id", "", "The read set's sequence store ID.")
	omics_startReadSetExportJobCmd.Flags().String("sources", "", "The job's source files.")
	omics_startReadSetExportJobCmd.MarkFlagRequired("destination")
	omics_startReadSetExportJobCmd.MarkFlagRequired("role-arn")
	omics_startReadSetExportJobCmd.MarkFlagRequired("sequence-store-id")
	omics_startReadSetExportJobCmd.MarkFlagRequired("sources")
	omicsCmd.AddCommand(omics_startReadSetExportJobCmd)
}
