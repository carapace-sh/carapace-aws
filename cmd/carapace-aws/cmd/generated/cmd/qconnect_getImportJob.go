package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getImportJobCmd = &cobra.Command{
	Use:   "get-import-job",
	Short: "Retrieves the started import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getImportJobCmd).Standalone()

		qconnect_getImportJobCmd.Flags().String("import-job-id", "", "The identifier of the import job to retrieve.")
		qconnect_getImportJobCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base that the import job belongs to.")
		qconnect_getImportJobCmd.MarkFlagRequired("import-job-id")
		qconnect_getImportJobCmd.MarkFlagRequired("knowledge-base-id")
	})
	qconnectCmd.AddCommand(qconnect_getImportJobCmd)
}
