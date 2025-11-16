package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getEvaluationJobCmd = &cobra.Command{
	Use:   "get-evaluation-job",
	Short: "Gets information about an evaluation job, such as the status of the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getEvaluationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getEvaluationJobCmd).Standalone()

		bedrock_getEvaluationJobCmd.Flags().String("job-identifier", "", "The Amazon Resource Name (ARN) of the evaluation job you want get information on.")
		bedrock_getEvaluationJobCmd.MarkFlagRequired("job-identifier")
	})
	bedrockCmd.AddCommand(bedrock_getEvaluationJobCmd)
}
