package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_invokeModelCmd = &cobra.Command{
	Use:   "invoke-model",
	Short: "Invokes the specified Amazon Bedrock model to run inference using the prompt and inference parameters provided in the request body.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_invokeModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockRuntime_invokeModelCmd).Standalone()

		bedrockRuntime_invokeModelCmd.Flags().String("accept", "", "The desired MIME type of the inference body in the response.")
		bedrockRuntime_invokeModelCmd.Flags().String("body", "", "The prompt and inference parameters in the format specified in the `contentType` in the header.")
		bedrockRuntime_invokeModelCmd.Flags().String("content-type", "", "The MIME type of the input data in the request.")
		bedrockRuntime_invokeModelCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail that you want to use.")
		bedrockRuntime_invokeModelCmd.Flags().String("guardrail-version", "", "The version number for the guardrail.")
		bedrockRuntime_invokeModelCmd.Flags().String("model-id", "", "The unique identifier of the model to invoke to run inference.")
		bedrockRuntime_invokeModelCmd.Flags().String("performance-config-latency", "", "Model performance settings for the request.")
		bedrockRuntime_invokeModelCmd.Flags().String("trace", "", "Specifies whether to enable or disable the Bedrock trace.")
		bedrockRuntime_invokeModelCmd.MarkFlagRequired("model-id")
	})
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_invokeModelCmd)
}
