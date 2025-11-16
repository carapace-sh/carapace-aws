package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invokes a Lambda function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_invokeCmd).Standalone()

	lambda_invokeCmd.Flags().String("client-context", "", "Up to 3,583 bytes of base64-encoded data about the invoking client to pass to the function in the context object.")
	lambda_invokeCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
	lambda_invokeCmd.Flags().String("invocation-type", "", "Choose from the following options.")
	lambda_invokeCmd.Flags().String("log-type", "", "Set to `Tail` to include the execution log in the response.")
	lambda_invokeCmd.Flags().String("payload", "", "The JSON that you want to provide to your Lambda function as input.")
	lambda_invokeCmd.Flags().String("qualifier", "", "Specify a version or alias to invoke a published version of the function.")
	lambda_invokeCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_invokeCmd)
}
