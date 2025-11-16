package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listBatchJobDefinitionsCmd = &cobra.Command{
	Use:   "list-batch-job-definitions",
	Short: "Lists all the available batch job definitions based on the batch job resources uploaded during the application creation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listBatchJobDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_listBatchJobDefinitionsCmd).Standalone()

		m2_listBatchJobDefinitionsCmd.Flags().String("application-id", "", "The identifier of the application.")
		m2_listBatchJobDefinitionsCmd.Flags().String("max-results", "", "The maximum number of batch job definitions to return.")
		m2_listBatchJobDefinitionsCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
		m2_listBatchJobDefinitionsCmd.Flags().String("prefix", "", "If the batch job definition is a FileBatchJobDefinition, the prefix allows you to search on the file names of FileBatchJobDefinitions.")
		m2_listBatchJobDefinitionsCmd.MarkFlagRequired("application-id")
	})
	m2Cmd.AddCommand(m2_listBatchJobDefinitionsCmd)
}
