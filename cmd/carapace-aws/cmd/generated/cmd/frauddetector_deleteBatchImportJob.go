package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteBatchImportJobCmd = &cobra.Command{
	Use:   "delete-batch-import-job",
	Short: "Deletes the specified batch import job ID record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteBatchImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteBatchImportJobCmd).Standalone()

		frauddetector_deleteBatchImportJobCmd.Flags().String("job-id", "", "The ID of the batch import job to delete.")
		frauddetector_deleteBatchImportJobCmd.MarkFlagRequired("job-id")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteBatchImportJobCmd)
}
