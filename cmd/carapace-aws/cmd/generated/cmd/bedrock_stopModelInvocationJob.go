package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_stopModelInvocationJobCmd = &cobra.Command{
	Use:   "stop-model-invocation-job",
	Short: "Stops a batch inference job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_stopModelInvocationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_stopModelInvocationJobCmd).Standalone()

		bedrock_stopModelInvocationJobCmd.Flags().String("job-identifier", "", "The Amazon Resource Name (ARN) of the batch inference job to stop.")
		bedrock_stopModelInvocationJobCmd.MarkFlagRequired("job-identifier")
	})
	bedrockCmd.AddCommand(bedrock_stopModelInvocationJobCmd)
}
