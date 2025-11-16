package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getModelInvocationJobCmd = &cobra.Command{
	Use:   "get-model-invocation-job",
	Short: "Gets details about a batch inference job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getModelInvocationJobCmd).Standalone()

	bedrock_getModelInvocationJobCmd.Flags().String("job-identifier", "", "The Amazon Resource Name (ARN) of the batch inference job.")
	bedrock_getModelInvocationJobCmd.MarkFlagRequired("job-identifier")
	bedrockCmd.AddCommand(bedrock_getModelInvocationJobCmd)
}
