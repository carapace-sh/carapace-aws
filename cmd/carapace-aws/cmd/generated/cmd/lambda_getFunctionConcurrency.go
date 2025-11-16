package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionConcurrencyCmd = &cobra.Command{
	Use:   "get-function-concurrency",
	Short: "Returns details about the reserved concurrency configuration for a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionConcurrencyCmd).Standalone()

	lambda_getFunctionConcurrencyCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_getFunctionConcurrencyCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionConcurrencyCmd)
}
