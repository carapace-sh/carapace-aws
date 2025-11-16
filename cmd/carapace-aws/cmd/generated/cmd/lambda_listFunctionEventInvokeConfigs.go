package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listFunctionEventInvokeConfigsCmd = &cobra.Command{
	Use:   "list-function-event-invoke-configs",
	Short: "Retrieves a list of configurations for asynchronous invocation for a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listFunctionEventInvokeConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_listFunctionEventInvokeConfigsCmd).Standalone()

		lambda_listFunctionEventInvokeConfigsCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_listFunctionEventInvokeConfigsCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
		lambda_listFunctionEventInvokeConfigsCmd.Flags().String("max-items", "", "The maximum number of configurations to return.")
		lambda_listFunctionEventInvokeConfigsCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_listFunctionEventInvokeConfigsCmd)
}
