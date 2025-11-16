package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listProvisionedConcurrencyConfigsCmd = &cobra.Command{
	Use:   "list-provisioned-concurrency-configs",
	Short: "Retrieves a list of provisioned concurrency configurations for a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listProvisionedConcurrencyConfigsCmd).Standalone()

	lambda_listProvisionedConcurrencyConfigsCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_listProvisionedConcurrencyConfigsCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
	lambda_listProvisionedConcurrencyConfigsCmd.Flags().String("max-items", "", "Specify a number to limit the number of configurations returned.")
	lambda_listProvisionedConcurrencyConfigsCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_listProvisionedConcurrencyConfigsCmd)
}
