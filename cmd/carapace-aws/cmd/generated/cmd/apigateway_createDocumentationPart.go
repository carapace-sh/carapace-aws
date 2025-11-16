package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createDocumentationPartCmd = &cobra.Command{
	Use:   "create-documentation-part",
	Short: "Creates a documentation part.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createDocumentationPartCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createDocumentationPartCmd).Standalone()

		apigateway_createDocumentationPartCmd.Flags().String("location", "", "The location of the targeted API entity of the to-be-created documentation part.")
		apigateway_createDocumentationPartCmd.Flags().String("properties", "", "The new documentation content map of the targeted API entity.")
		apigateway_createDocumentationPartCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_createDocumentationPartCmd.MarkFlagRequired("location")
		apigateway_createDocumentationPartCmd.MarkFlagRequired("properties")
		apigateway_createDocumentationPartCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_createDocumentationPartCmd)
}
