package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_getAsyncInvokeCmd = &cobra.Command{
	Use:   "get-async-invoke",
	Short: "Retrieve information about an asynchronous invocation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_getAsyncInvokeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockRuntime_getAsyncInvokeCmd).Standalone()

		bedrockRuntime_getAsyncInvokeCmd.Flags().String("invocation-arn", "", "The invocation's ARN.")
		bedrockRuntime_getAsyncInvokeCmd.MarkFlagRequired("invocation-arn")
	})
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_getAsyncInvokeCmd)
}
