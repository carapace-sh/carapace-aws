package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_putProvisionedConcurrencyConfigCmd = &cobra.Command{
	Use:   "put-provisioned-concurrency-config",
	Short: "Adds a provisioned concurrency configuration to a function's alias or version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_putProvisionedConcurrencyConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_putProvisionedConcurrencyConfigCmd).Standalone()

		lambda_putProvisionedConcurrencyConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_putProvisionedConcurrencyConfigCmd.Flags().String("provisioned-concurrent-executions", "", "The amount of provisioned concurrency to allocate for the version or alias.")
		lambda_putProvisionedConcurrencyConfigCmd.Flags().String("qualifier", "", "The version number or alias name.")
		lambda_putProvisionedConcurrencyConfigCmd.MarkFlagRequired("function-name")
		lambda_putProvisionedConcurrencyConfigCmd.MarkFlagRequired("provisioned-concurrent-executions")
		lambda_putProvisionedConcurrencyConfigCmd.MarkFlagRequired("qualifier")
	})
	lambdaCmd.AddCommand(lambda_putProvisionedConcurrencyConfigCmd)
}
