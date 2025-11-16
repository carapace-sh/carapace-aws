package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteFunctionCodeSigningConfigCmd = &cobra.Command{
	Use:   "delete-function-code-signing-config",
	Short: "Removes the code signing configuration from the function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteFunctionCodeSigningConfigCmd).Standalone()

	lambda_deleteFunctionCodeSigningConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_deleteFunctionCodeSigningConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_deleteFunctionCodeSigningConfigCmd)
}
