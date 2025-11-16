package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getModelCopyJobCmd = &cobra.Command{
	Use:   "get-model-copy-job",
	Short: "Retrieves information about a model copy job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getModelCopyJobCmd).Standalone()

	bedrock_getModelCopyJobCmd.Flags().String("job-arn", "", "The Amazon Resource Name (ARN) of the model copy job.")
	bedrock_getModelCopyJobCmd.MarkFlagRequired("job-arn")
	bedrockCmd.AddCommand(bedrock_getModelCopyJobCmd)
}
