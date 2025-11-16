package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDocumentationPartCmd = &cobra.Command{
	Use:   "get-documentation-part",
	Short: "Gets a documentation part.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDocumentationPartCmd).Standalone()

	apigateway_getDocumentationPartCmd.Flags().String("documentation-part-id", "", "The string identifier of the associated RestApi.")
	apigateway_getDocumentationPartCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getDocumentationPartCmd.MarkFlagRequired("documentation-part-id")
	apigateway_getDocumentationPartCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getDocumentationPartCmd)
}
