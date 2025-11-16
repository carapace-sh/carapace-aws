package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_putFunctionEventInvokeConfigCmd = &cobra.Command{
	Use:   "put-function-event-invoke-config",
	Short: "Configures options for [asynchronous invocation](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html) on a function, version, or alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_putFunctionEventInvokeConfigCmd).Standalone()

	lambda_putFunctionEventInvokeConfigCmd.Flags().String("destination-config", "", "A destination for events after they have been sent to a function for processing.")
	lambda_putFunctionEventInvokeConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
	lambda_putFunctionEventInvokeConfigCmd.Flags().String("maximum-event-age-in-seconds", "", "The maximum age of a request that Lambda sends to a function for processing.")
	lambda_putFunctionEventInvokeConfigCmd.Flags().String("maximum-retry-attempts", "", "The maximum number of times to retry when the function returns an error.")
	lambda_putFunctionEventInvokeConfigCmd.Flags().String("qualifier", "", "A version number or alias name.")
	lambda_putFunctionEventInvokeConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_putFunctionEventInvokeConfigCmd)
}
