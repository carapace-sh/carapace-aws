package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createInferenceProfileCmd = &cobra.Command{
	Use:   "create-inference-profile",
	Short: "Creates an application inference profile to track metrics and costs when invoking a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createInferenceProfileCmd).Standalone()

	bedrock_createInferenceProfileCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrock_createInferenceProfileCmd.Flags().String("description", "", "A description for the inference profile.")
	bedrock_createInferenceProfileCmd.Flags().String("inference-profile-name", "", "A name for the inference profile.")
	bedrock_createInferenceProfileCmd.Flags().String("model-source", "", "The foundation model or system-defined inference profile that the inference profile will track metrics and costs for.")
	bedrock_createInferenceProfileCmd.Flags().String("tags", "", "An array of objects, each of which contains a tag and its value.")
	bedrock_createInferenceProfileCmd.MarkFlagRequired("inference-profile-name")
	bedrock_createInferenceProfileCmd.MarkFlagRequired("model-source")
	bedrockCmd.AddCommand(bedrock_createInferenceProfileCmd)
}
