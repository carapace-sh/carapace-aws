package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateFunctionUrlConfigCmd = &cobra.Command{
	Use:   "update-function-url-config",
	Short: "Updates the configuration for a Lambda function URL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateFunctionUrlConfigCmd).Standalone()

	lambda_updateFunctionUrlConfigCmd.Flags().String("auth-type", "", "The type of authentication that your function URL uses.")
	lambda_updateFunctionUrlConfigCmd.Flags().String("cors", "", "The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.")
	lambda_updateFunctionUrlConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_updateFunctionUrlConfigCmd.Flags().String("invoke-mode", "", "Use one of the following options:")
	lambda_updateFunctionUrlConfigCmd.Flags().String("qualifier", "", "The alias name.")
	lambda_updateFunctionUrlConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_updateFunctionUrlConfigCmd)
}
