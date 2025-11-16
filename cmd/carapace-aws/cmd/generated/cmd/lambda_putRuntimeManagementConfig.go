package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_putRuntimeManagementConfigCmd = &cobra.Command{
	Use:   "put-runtime-management-config",
	Short: "Sets the runtime management configuration for a function's version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_putRuntimeManagementConfigCmd).Standalone()

	lambda_putRuntimeManagementConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_putRuntimeManagementConfigCmd.Flags().String("qualifier", "", "Specify a version of the function.")
	lambda_putRuntimeManagementConfigCmd.Flags().String("runtime-version-arn", "", "The ARN of the runtime version you want the function to use.")
	lambda_putRuntimeManagementConfigCmd.Flags().String("update-runtime-on", "", "Specify the runtime update mode.")
	lambda_putRuntimeManagementConfigCmd.MarkFlagRequired("function-name")
	lambda_putRuntimeManagementConfigCmd.MarkFlagRequired("update-runtime-on")
	lambdaCmd.AddCommand(lambda_putRuntimeManagementConfigCmd)
}
