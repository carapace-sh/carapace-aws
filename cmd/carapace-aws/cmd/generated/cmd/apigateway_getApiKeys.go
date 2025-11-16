package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getApiKeysCmd = &cobra.Command{
	Use:   "get-api-keys",
	Short: "Gets information about the current ApiKeys resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getApiKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getApiKeysCmd).Standalone()

		apigateway_getApiKeysCmd.Flags().String("customer-id", "", "The identifier of a customer in Amazon Web Services Marketplace or an external system, such as a developer portal.")
		apigateway_getApiKeysCmd.Flags().String("include-values", "", "A boolean flag to specify whether (`true`) or not (`false`) the result contains key values.")
		apigateway_getApiKeysCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getApiKeysCmd.Flags().String("name-query", "", "The name of queried API keys.")
		apigateway_getApiKeysCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	})
	apigatewayCmd.AddCommand(apigateway_getApiKeysCmd)
}
