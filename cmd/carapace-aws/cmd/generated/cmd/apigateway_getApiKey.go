package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getApiKeyCmd = &cobra.Command{
	Use:   "get-api-key",
	Short: "Gets information about the current ApiKey resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getApiKeyCmd).Standalone()

	apigateway_getApiKeyCmd.Flags().String("api-key", "", "The identifier of the ApiKey resource.")
	apigateway_getApiKeyCmd.Flags().String("include-value", "", "A boolean flag to specify whether (`true`) or not (`false`) the result contains the key value.")
	apigateway_getApiKeyCmd.MarkFlagRequired("api-key")
	apigatewayCmd.AddCommand(apigateway_getApiKeyCmd)
}
