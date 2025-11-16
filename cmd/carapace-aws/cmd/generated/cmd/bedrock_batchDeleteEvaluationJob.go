package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_batchDeleteEvaluationJobCmd = &cobra.Command{
	Use:   "batch-delete-evaluation-job",
	Short: "Deletes a batch of evaluation jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_batchDeleteEvaluationJobCmd).Standalone()

	bedrock_batchDeleteEvaluationJobCmd.Flags().String("job-identifiers", "", "A list of one or more evaluation job Amazon Resource Names (ARNs) you want to delete.")
	bedrock_batchDeleteEvaluationJobCmd.MarkFlagRequired("job-identifiers")
	bedrockCmd.AddCommand(bedrock_batchDeleteEvaluationJobCmd)
}
