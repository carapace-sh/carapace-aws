package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionCodeSigningConfigCmd = &cobra.Command{
	Use:   "get-function-code-signing-config",
	Short: "Returns the code signing configuration for the specified function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionCodeSigningConfigCmd).Standalone()

	lambda_getFunctionCodeSigningConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_getFunctionCodeSigningConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionCodeSigningConfigCmd)
}
