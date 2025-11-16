package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateFunctionEventInvokeConfigCmd = &cobra.Command{
	Use:   "update-function-event-invoke-config",
	Short: "Updates the configuration for asynchronous invocation for a function, version, or alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateFunctionEventInvokeConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_updateFunctionEventInvokeConfigCmd).Standalone()

		lambda_updateFunctionEventInvokeConfigCmd.Flags().String("destination-config", "", "A destination for events after they have been sent to a function for processing.")
		lambda_updateFunctionEventInvokeConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
		lambda_updateFunctionEventInvokeConfigCmd.Flags().String("maximum-event-age-in-seconds", "", "The maximum age of a request that Lambda sends to a function for processing.")
		lambda_updateFunctionEventInvokeConfigCmd.Flags().String("maximum-retry-attempts", "", "The maximum number of times to retry when the function returns an error.")
		lambda_updateFunctionEventInvokeConfigCmd.Flags().String("qualifier", "", "A version number or alias name.")
		lambda_updateFunctionEventInvokeConfigCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_updateFunctionEventInvokeConfigCmd)
}
