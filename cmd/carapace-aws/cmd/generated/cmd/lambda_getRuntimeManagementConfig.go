package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getRuntimeManagementConfigCmd = &cobra.Command{
	Use:   "get-runtime-management-config",
	Short: "Retrieves the runtime management configuration for a function's version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getRuntimeManagementConfigCmd).Standalone()

	lambda_getRuntimeManagementConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_getRuntimeManagementConfigCmd.Flags().String("qualifier", "", "Specify a version of the function.")
	lambda_getRuntimeManagementConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getRuntimeManagementConfigCmd)
}
