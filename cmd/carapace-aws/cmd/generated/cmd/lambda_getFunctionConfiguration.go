package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionConfigurationCmd = &cobra.Command{
	Use:   "get-function-configuration",
	Short: "Returns the version-specific settings of a Lambda function or version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionConfigurationCmd).Standalone()

	lambda_getFunctionConfigurationCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
	lambda_getFunctionConfigurationCmd.Flags().String("qualifier", "", "Specify a version or alias to get details about a published version of the function.")
	lambda_getFunctionConfigurationCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionConfigurationCmd)
}
