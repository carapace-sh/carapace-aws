package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_cancelBatchImportJobCmd = &cobra.Command{
	Use:   "cancel-batch-import-job",
	Short: "Cancels an in-progress batch import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_cancelBatchImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_cancelBatchImportJobCmd).Standalone()

		frauddetector_cancelBatchImportJobCmd.Flags().String("job-id", "", "The ID of an in-progress batch import job to cancel.")
		frauddetector_cancelBatchImportJobCmd.MarkFlagRequired("job-id")
	})
	frauddetectorCmd.AddCommand(frauddetector_cancelBatchImportJobCmd)
}
