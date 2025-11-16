package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listFoundationModelsCmd = &cobra.Command{
	Use:   "list-foundation-models",
	Short: "Lists Amazon Bedrock foundation models that you can use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listFoundationModelsCmd).Standalone()

	bedrock_listFoundationModelsCmd.Flags().String("by-customization-type", "", "Return models that support the customization type that you specify.")
	bedrock_listFoundationModelsCmd.Flags().String("by-inference-type", "", "Return models that support the inference type that you specify.")
	bedrock_listFoundationModelsCmd.Flags().String("by-output-modality", "", "Return models that support the output modality that you specify.")
	bedrock_listFoundationModelsCmd.Flags().String("by-provider", "", "Return models belonging to the model provider that you specify.")
	bedrockCmd.AddCommand(bedrock_listFoundationModelsCmd)
}
