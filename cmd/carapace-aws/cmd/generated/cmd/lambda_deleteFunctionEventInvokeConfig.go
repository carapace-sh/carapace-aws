package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteFunctionEventInvokeConfigCmd = &cobra.Command{
	Use:   "delete-function-event-invoke-config",
	Short: "Deletes the configuration for asynchronous invocation for a function, version, or alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteFunctionEventInvokeConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_deleteFunctionEventInvokeConfigCmd).Standalone()

		lambda_deleteFunctionEventInvokeConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
		lambda_deleteFunctionEventInvokeConfigCmd.Flags().String("qualifier", "", "A version number or alias name.")
		lambda_deleteFunctionEventInvokeConfigCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_deleteFunctionEventInvokeConfigCmd)
}
