package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_deleteImportJobCmd = &cobra.Command{
	Use:   "delete-import-job",
	Short: "Deletes the quick response import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_deleteImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_deleteImportJobCmd).Standalone()

		wisdom_deleteImportJobCmd.Flags().String("import-job-id", "", "The identifier of the import job to be deleted.")
		wisdom_deleteImportJobCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_deleteImportJobCmd.MarkFlagRequired("import-job-id")
		wisdom_deleteImportJobCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_deleteImportJobCmd)
}
