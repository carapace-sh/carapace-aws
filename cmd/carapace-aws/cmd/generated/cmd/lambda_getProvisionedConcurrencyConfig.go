package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getProvisionedConcurrencyConfigCmd = &cobra.Command{
	Use:   "get-provisioned-concurrency-config",
	Short: "Retrieves the provisioned concurrency configuration for a function's alias or version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getProvisionedConcurrencyConfigCmd).Standalone()

	lambda_getProvisionedConcurrencyConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_getProvisionedConcurrencyConfigCmd.Flags().String("qualifier", "", "The version number or alias name.")
	lambda_getProvisionedConcurrencyConfigCmd.MarkFlagRequired("function-name")
	lambda_getProvisionedConcurrencyConfigCmd.MarkFlagRequired("qualifier")
	lambdaCmd.AddCommand(lambda_getProvisionedConcurrencyConfigCmd)
}
