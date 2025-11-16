package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getModelImportJobCmd = &cobra.Command{
	Use:   "get-model-import-job",
	Short: "Retrieves the properties associated with import model job, including the status of the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getModelImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getModelImportJobCmd).Standalone()

		bedrock_getModelImportJobCmd.Flags().String("job-identifier", "", "The identifier of the import job.")
		bedrock_getModelImportJobCmd.MarkFlagRequired("job-identifier")
	})
	bedrockCmd.AddCommand(bedrock_getModelImportJobCmd)
}
