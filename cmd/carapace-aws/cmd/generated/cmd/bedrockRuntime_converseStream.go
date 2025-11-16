package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_converseStreamCmd = &cobra.Command{
	Use:   "converse-stream",
	Short: "Sends messages to the specified Amazon Bedrock model and returns the response in a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_converseStreamCmd).Standalone()

	bedrockRuntime_converseStreamCmd.Flags().String("additional-model-request-fields", "", "Additional inference parameters that the model supports, beyond the base set of inference parameters that `Converse` and `ConverseStream` support in the `inferenceConfig` field.")
	bedrockRuntime_converseStreamCmd.Flags().String("additional-model-response-field-paths", "", "Additional model parameters field paths to return in the response.")
	bedrockRuntime_converseStreamCmd.Flags().String("guardrail-config", "", "Configuration information for a guardrail that you want to use in the request.")
	bedrockRuntime_converseStreamCmd.Flags().String("inference-config", "", "Inference parameters to pass to the model.")
	bedrockRuntime_converseStreamCmd.Flags().String("messages", "", "The messages that you want to send to the model.")
	bedrockRuntime_converseStreamCmd.Flags().String("model-id", "", "Specifies the model or throughput with which to run inference, or the prompt resource to use in inference.")
	bedrockRuntime_converseStreamCmd.Flags().String("performance-config", "", "Model performance settings for the request.")
	bedrockRuntime_converseStreamCmd.Flags().String("prompt-variables", "", "Contains a map of variables in a prompt from Prompt management to objects containing the values to fill in for them when running model invocation.")
	bedrockRuntime_converseStreamCmd.Flags().String("request-metadata", "", "Key-value pairs that you can use to filter invocation logs.")
	bedrockRuntime_converseStreamCmd.Flags().String("system", "", "A prompt that provides instructions or context to the model about the task it should perform, or the persona it should adopt during the conversation.")
	bedrockRuntime_converseStreamCmd.Flags().String("tool-config", "", "Configuration information for the tools that the model can use when generating a response.")
	bedrockRuntime_converseStreamCmd.MarkFlagRequired("model-id")
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_converseStreamCmd)
}
