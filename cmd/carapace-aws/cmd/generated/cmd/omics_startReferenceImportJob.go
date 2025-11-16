package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_startReferenceImportJobCmd = &cobra.Command{
	Use:   "start-reference-import-job",
	Short: "Imports a reference genome from Amazon S3 into a specified reference store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_startReferenceImportJobCmd).Standalone()

	omics_startReferenceImportJobCmd.Flags().String("client-token", "", "To ensure that jobs don't run multiple times, specify a unique token for each job.")
	omics_startReferenceImportJobCmd.Flags().String("reference-store-id", "", "The job's reference store ID.")
	omics_startReferenceImportJobCmd.Flags().String("role-arn", "", "A service role for the job.")
	omics_startReferenceImportJobCmd.Flags().String("sources", "", "The job's source files.")
	omics_startReferenceImportJobCmd.MarkFlagRequired("reference-store-id")
	omics_startReferenceImportJobCmd.MarkFlagRequired("role-arn")
	omics_startReferenceImportJobCmd.MarkFlagRequired("sources")
	omicsCmd.AddCommand(omics_startReferenceImportJobCmd)
}
