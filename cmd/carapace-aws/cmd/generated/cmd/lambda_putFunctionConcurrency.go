package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_putFunctionConcurrencyCmd = &cobra.Command{
	Use:   "put-function-concurrency",
	Short: "Sets the maximum number of simultaneous executions for a function, and reserves capacity for that concurrency level.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_putFunctionConcurrencyCmd).Standalone()

	lambda_putFunctionConcurrencyCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_putFunctionConcurrencyCmd.Flags().String("reserved-concurrent-executions", "", "The number of simultaneous executions to reserve for the function.")
	lambda_putFunctionConcurrencyCmd.MarkFlagRequired("function-name")
	lambda_putFunctionConcurrencyCmd.MarkFlagRequired("reserved-concurrent-executions")
	lambdaCmd.AddCommand(lambda_putFunctionConcurrencyCmd)
}
