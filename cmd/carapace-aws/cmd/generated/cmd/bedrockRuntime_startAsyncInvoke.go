package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_startAsyncInvokeCmd = &cobra.Command{
	Use:   "start-async-invoke",
	Short: "Starts an asynchronous invocation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_startAsyncInvokeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockRuntime_startAsyncInvokeCmd).Standalone()

		bedrockRuntime_startAsyncInvokeCmd.Flags().String("client-request-token", "", "Specify idempotency token to ensure that requests are not duplicated.")
		bedrockRuntime_startAsyncInvokeCmd.Flags().String("model-id", "", "The model to invoke.")
		bedrockRuntime_startAsyncInvokeCmd.Flags().String("model-input", "", "Input to send to the model.")
		bedrockRuntime_startAsyncInvokeCmd.Flags().String("output-data-config", "", "Where to store the output.")
		bedrockRuntime_startAsyncInvokeCmd.Flags().String("tags", "", "Tags to apply to the invocation.")
		bedrockRuntime_startAsyncInvokeCmd.MarkFlagRequired("model-id")
		bedrockRuntime_startAsyncInvokeCmd.MarkFlagRequired("model-input")
		bedrockRuntime_startAsyncInvokeCmd.MarkFlagRequired("output-data-config")
	})
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_startAsyncInvokeCmd)
}
