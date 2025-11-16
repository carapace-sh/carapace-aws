package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_converseCmd = &cobra.Command{
	Use:   "converse",
	Short: "Sends messages to the specified Amazon Bedrock model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_converseCmd).Standalone()

	bedrockRuntime_converseCmd.Flags().String("additional-model-request-fields", "", "Additional inference parameters that the model supports, beyond the base set of inference parameters that `Converse` and `ConverseStream` support in the `inferenceConfig` field.")
	bedrockRuntime_converseCmd.Flags().String("additional-model-response-field-paths", "", "Additional model parameters field paths to return in the response.")
	bedrockRuntime_converseCmd.Flags().String("guardrail-config", "", "Configuration information for a guardrail that you want to use in the request.")
	bedrockRuntime_converseCmd.Flags().String("inference-config", "", "Inference parameters to pass to the model.")
	bedrockRuntime_converseCmd.Flags().String("messages", "", "The messages that you want to send to the model.")
	bedrockRuntime_converseCmd.Flags().String("model-id", "", "Specifies the model or throughput with which to run inference, or the prompt resource to use in inference.")
	bedrockRuntime_converseCmd.Flags().String("performance-config", "", "Model performance settings for the request.")
	bedrockRuntime_converseCmd.Flags().String("prompt-variables", "", "Contains a map of variables in a prompt from Prompt management to objects containing the values to fill in for them when running model invocation.")
	bedrockRuntime_converseCmd.Flags().String("request-metadata", "", "Key-value pairs that you can use to filter invocation logs.")
	bedrockRuntime_converseCmd.Flags().String("system", "", "A prompt that provides instructions or context to the model about the task it should perform, or the persona it should adopt during the conversation.")
	bedrockRuntime_converseCmd.Flags().String("tool-config", "", "Configuration information for the tools that the model can use when generating a response.")
	bedrockRuntime_converseCmd.MarkFlagRequired("model-id")
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_converseCmd)
}
