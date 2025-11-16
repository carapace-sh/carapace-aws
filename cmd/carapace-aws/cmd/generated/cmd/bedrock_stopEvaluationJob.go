package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_stopEvaluationJobCmd = &cobra.Command{
	Use:   "stop-evaluation-job",
	Short: "Stops an evaluation job that is current being created or running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_stopEvaluationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_stopEvaluationJobCmd).Standalone()

		bedrock_stopEvaluationJobCmd.Flags().String("job-identifier", "", "The Amazon Resource Name (ARN) of the evaluation job you want to stop.")
		bedrock_stopEvaluationJobCmd.MarkFlagRequired("job-identifier")
	})
	bedrockCmd.AddCommand(bedrock_stopEvaluationJobCmd)
}
