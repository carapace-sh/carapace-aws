package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteDocumentationPartCmd = &cobra.Command{
	Use:   "delete-documentation-part",
	Short: "Deletes a documentation part",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteDocumentationPartCmd).Standalone()

	apigateway_deleteDocumentationPartCmd.Flags().String("documentation-part-id", "", "The identifier of the to-be-deleted documentation part.")
	apigateway_deleteDocumentationPartCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_deleteDocumentationPartCmd.MarkFlagRequired("documentation-part-id")
	apigateway_deleteDocumentationPartCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_deleteDocumentationPartCmd)
}
