package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getResourcesCmd = &cobra.Command{
	Use:   "get-resources",
	Short: "Lists information about a collection of Resource resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getResourcesCmd).Standalone()

	apigateway_getResourcesCmd.Flags().String("embed", "", "A query parameter used to retrieve the specified resources embedded in the returned Resources resource in the response.")
	apigateway_getResourcesCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getResourcesCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigateway_getResourcesCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getResourcesCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getResourcesCmd)
}
