package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteImportJobCmd = &cobra.Command{
	Use:   "delete-import-job",
	Short: "Deletes the quick response import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteImportJobCmd).Standalone()

	qconnect_deleteImportJobCmd.Flags().String("import-job-id", "", "The identifier of the import job to be deleted.")
	qconnect_deleteImportJobCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_deleteImportJobCmd.MarkFlagRequired("import-job-id")
	qconnect_deleteImportJobCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_deleteImportJobCmd)
}
