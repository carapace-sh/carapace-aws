package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_invokeAsyncCmd = &cobra.Command{
	Use:   "invoke-async",
	Short: "For asynchronous function invocation, use [Invoke]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_invokeAsyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_invokeAsyncCmd).Standalone()

		lambda_invokeAsyncCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_invokeAsyncCmd.Flags().String("invoke-args", "", "The JSON that you want to provide to your Lambda function as input.")
		lambda_invokeAsyncCmd.MarkFlagRequired("function-name")
		lambda_invokeAsyncCmd.MarkFlagRequired("invoke-args")
	})
	lambdaCmd.AddCommand(lambda_invokeAsyncCmd)
}
