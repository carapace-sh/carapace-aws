package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_invokeWithResponseStreamCmd = &cobra.Command{
	Use:   "invoke-with-response-stream",
	Short: "Configure your Lambda functions to stream response payloads back to clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_invokeWithResponseStreamCmd).Standalone()

	lambda_invokeWithResponseStreamCmd.Flags().String("client-context", "", "Up to 3,583 bytes of base64-encoded data about the invoking client to pass to the function in the context object.")
	lambda_invokeWithResponseStreamCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_invokeWithResponseStreamCmd.Flags().String("invocation-type", "", "Use one of the following options:")
	lambda_invokeWithResponseStreamCmd.Flags().String("log-type", "", "Set to `Tail` to include the execution log in the response.")
	lambda_invokeWithResponseStreamCmd.Flags().String("payload", "", "The JSON that you want to provide to your Lambda function as input.")
	lambda_invokeWithResponseStreamCmd.Flags().String("qualifier", "", "The alias name.")
	lambda_invokeWithResponseStreamCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_invokeWithResponseStreamCmd)
}
