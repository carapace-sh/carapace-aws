package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionEventInvokeConfigCmd = &cobra.Command{
	Use:   "get-function-event-invoke-config",
	Short: "Retrieves the configuration for asynchronous invocation for a function, version, or alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionEventInvokeConfigCmd).Standalone()

	lambda_getFunctionEventInvokeConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
	lambda_getFunctionEventInvokeConfigCmd.Flags().String("qualifier", "", "A version number or alias name.")
	lambda_getFunctionEventInvokeConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionEventInvokeConfigCmd)
}
