package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getResourceCmd = &cobra.Command{
	Use:   "get-resource",
	Short: "Lists information about a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getResourceCmd).Standalone()

	apigateway_getResourceCmd.Flags().String("embed", "", "A query parameter to retrieve the specified resources embedded in the returned Resource representation in the response.")
	apigateway_getResourceCmd.Flags().String("resource-id", "", "The identifier for the Resource resource.")
	apigateway_getResourceCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getResourceCmd.MarkFlagRequired("resource-id")
	apigateway_getResourceCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getResourceCmd)
}
