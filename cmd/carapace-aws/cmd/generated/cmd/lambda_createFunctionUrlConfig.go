package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_createFunctionUrlConfigCmd = &cobra.Command{
	Use:   "create-function-url-config",
	Short: "Creates a Lambda function URL with the specified configuration parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_createFunctionUrlConfigCmd).Standalone()

	lambda_createFunctionUrlConfigCmd.Flags().String("auth-type", "", "The type of authentication that your function URL uses.")
	lambda_createFunctionUrlConfigCmd.Flags().String("cors", "", "The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.")
	lambda_createFunctionUrlConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_createFunctionUrlConfigCmd.Flags().String("invoke-mode", "", "Use one of the following options:")
	lambda_createFunctionUrlConfigCmd.Flags().String("qualifier", "", "The alias name.")
	lambda_createFunctionUrlConfigCmd.MarkFlagRequired("auth-type")
	lambda_createFunctionUrlConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_createFunctionUrlConfigCmd)
}
