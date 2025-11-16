package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_putFunctionCodeSigningConfigCmd = &cobra.Command{
	Use:   "put-function-code-signing-config",
	Short: "Update the code signing configuration for the function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_putFunctionCodeSigningConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_putFunctionCodeSigningConfigCmd).Standalone()

		lambda_putFunctionCodeSigningConfigCmd.Flags().String("code-signing-config-arn", "", "The The Amazon Resource Name (ARN) of the code signing configuration.")
		lambda_putFunctionCodeSigningConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_putFunctionCodeSigningConfigCmd.MarkFlagRequired("code-signing-config-arn")
		lambda_putFunctionCodeSigningConfigCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_putFunctionCodeSigningConfigCmd)
}
