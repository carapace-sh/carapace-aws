package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_invokeModelWithBidirectionalStreamCmd = &cobra.Command{
	Use:   "invoke-model-with-bidirectional-stream",
	Short: "Invoke the specified Amazon Bedrock model to run inference using the bidirectional stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_invokeModelWithBidirectionalStreamCmd).Standalone()

	bedrockRuntime_invokeModelWithBidirectionalStreamCmd.Flags().String("body", "", "The prompt and inference parameters in the format specified in the `BidirectionalInputPayloadPart` in the header.")
	bedrockRuntime_invokeModelWithBidirectionalStreamCmd.Flags().String("model-id", "", "The model ID or ARN of the model ID to use.")
	bedrockRuntime_invokeModelWithBidirectionalStreamCmd.MarkFlagRequired("body")
	bedrockRuntime_invokeModelWithBidirectionalStreamCmd.MarkFlagRequired("model-id")
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_invokeModelWithBidirectionalStreamCmd)
}
