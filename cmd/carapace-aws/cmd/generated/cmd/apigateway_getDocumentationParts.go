package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDocumentationPartsCmd = &cobra.Command{
	Use:   "get-documentation-parts",
	Short: "Gets documentation parts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDocumentationPartsCmd).Standalone()

	apigateway_getDocumentationPartsCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getDocumentationPartsCmd.Flags().String("location-status", "", "The status of the API documentation parts to retrieve.")
	apigateway_getDocumentationPartsCmd.Flags().String("name-query", "", "The name of API entities of the to-be-retrieved documentation parts.")
	apigateway_getDocumentationPartsCmd.Flags().String("path", "", "The path of API entities of the to-be-retrieved documentation parts.")
	apigateway_getDocumentationPartsCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigateway_getDocumentationPartsCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getDocumentationPartsCmd.Flags().String("type", "", "The type of API entities of the to-be-retrieved documentation parts.")
	apigateway_getDocumentationPartsCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getDocumentationPartsCmd)
}
