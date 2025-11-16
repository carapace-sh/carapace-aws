package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteDocumentationVersionCmd = &cobra.Command{
	Use:   "delete-documentation-version",
	Short: "Deletes a documentation version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteDocumentationVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteDocumentationVersionCmd).Standalone()

		apigateway_deleteDocumentationVersionCmd.Flags().String("documentation-version", "", "The version identifier of a to-be-deleted documentation snapshot.")
		apigateway_deleteDocumentationVersionCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteDocumentationVersionCmd.MarkFlagRequired("documentation-version")
		apigateway_deleteDocumentationVersionCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteDocumentationVersionCmd)
}
