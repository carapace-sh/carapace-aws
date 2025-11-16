package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteProvisionedConcurrencyConfigCmd = &cobra.Command{
	Use:   "delete-provisioned-concurrency-config",
	Short: "Deletes the provisioned concurrency configuration for a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteProvisionedConcurrencyConfigCmd).Standalone()

	lambda_deleteProvisionedConcurrencyConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_deleteProvisionedConcurrencyConfigCmd.Flags().String("qualifier", "", "The version number or alias name.")
	lambda_deleteProvisionedConcurrencyConfigCmd.MarkFlagRequired("function-name")
	lambda_deleteProvisionedConcurrencyConfigCmd.MarkFlagRequired("qualifier")
	lambdaCmd.AddCommand(lambda_deleteProvisionedConcurrencyConfigCmd)
}
