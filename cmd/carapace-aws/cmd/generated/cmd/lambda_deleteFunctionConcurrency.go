package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteFunctionConcurrencyCmd = &cobra.Command{
	Use:   "delete-function-concurrency",
	Short: "Removes a concurrent execution limit from a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteFunctionConcurrencyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_deleteFunctionConcurrencyCmd).Standalone()

		lambda_deleteFunctionConcurrencyCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_deleteFunctionConcurrencyCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_deleteFunctionConcurrencyCmd)
}
