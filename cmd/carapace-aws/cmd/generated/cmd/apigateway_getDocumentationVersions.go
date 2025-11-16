package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDocumentationVersionsCmd = &cobra.Command{
	Use:   "get-documentation-versions",
	Short: "Gets documentation versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDocumentationVersionsCmd).Standalone()

	apigateway_getDocumentationVersionsCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getDocumentationVersionsCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigateway_getDocumentationVersionsCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getDocumentationVersionsCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getDocumentationVersionsCmd)
}
