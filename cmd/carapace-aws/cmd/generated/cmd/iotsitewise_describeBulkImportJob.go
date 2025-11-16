package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeBulkImportJobCmd = &cobra.Command{
	Use:   "describe-bulk-import-job",
	Short: "Retrieves information about a bulk import job request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeBulkImportJobCmd).Standalone()

	iotsitewise_describeBulkImportJobCmd.Flags().String("job-id", "", "The ID of the job.")
	iotsitewise_describeBulkImportJobCmd.MarkFlagRequired("job-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeBulkImportJobCmd)
}
