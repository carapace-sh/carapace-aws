package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_invokeModelWithResponseStreamCmd = &cobra.Command{
	Use:   "invoke-model-with-response-stream",
	Short: "Invoke the specified Amazon Bedrock model to run inference using the prompt and inference parameters provided in the request body.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_invokeModelWithResponseStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockRuntime_invokeModelWithResponseStreamCmd).Standalone()

		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("accept", "", "The desired MIME type of the inference body in the response.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("body", "", "The prompt and inference parameters in the format specified in the `contentType` in the header.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("content-type", "", "The MIME type of the input data in the request.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail that you want to use.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("guardrail-version", "", "The version number for the guardrail.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("model-id", "", "The unique identifier of the model to invoke to run inference.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("performance-config-latency", "", "Model performance settings for the request.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.Flags().String("trace", "", "Specifies whether to enable or disable the Bedrock trace.")
		bedrockRuntime_invokeModelWithResponseStreamCmd.MarkFlagRequired("model-id")
	})
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_invokeModelWithResponseStreamCmd)
}
