package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_putFunctionRecursionConfigCmd = &cobra.Command{
	Use:   "put-function-recursion-config",
	Short: "Sets your function's [recursive loop detection](https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html) configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_putFunctionRecursionConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_putFunctionRecursionConfigCmd).Standalone()

		lambda_putFunctionRecursionConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_putFunctionRecursionConfigCmd.Flags().String("recursive-loop", "", "If you set your function's recursive loop detection configuration to `Allow`, Lambda doesn't take any action when it detects your function being invoked as part of a recursive loop.")
		lambda_putFunctionRecursionConfigCmd.MarkFlagRequired("function-name")
		lambda_putFunctionRecursionConfigCmd.MarkFlagRequired("recursive-loop")
	})
	lambdaCmd.AddCommand(lambda_putFunctionRecursionConfigCmd)
}
