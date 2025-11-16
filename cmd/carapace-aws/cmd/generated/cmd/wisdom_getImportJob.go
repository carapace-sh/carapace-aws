package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getImportJobCmd = &cobra.Command{
	Use:   "get-import-job",
	Short: "Retrieves the started import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getImportJobCmd).Standalone()

	wisdom_getImportJobCmd.Flags().String("import-job-id", "", "The identifier of the import job to retrieve.")
	wisdom_getImportJobCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base that the import job belongs to.")
	wisdom_getImportJobCmd.MarkFlagRequired("import-job-id")
	wisdom_getImportJobCmd.MarkFlagRequired("knowledge-base-id")
	wisdomCmd.AddCommand(wisdom_getImportJobCmd)
}
