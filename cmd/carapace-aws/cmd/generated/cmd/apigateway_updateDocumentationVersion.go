package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateDocumentationVersionCmd = &cobra.Command{
	Use:   "update-documentation-version",
	Short: "Updates a documentation version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateDocumentationVersionCmd).Standalone()

	apigateway_updateDocumentationVersionCmd.Flags().String("documentation-version", "", "The version identifier of the to-be-updated documentation version.")
	apigateway_updateDocumentationVersionCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateDocumentationVersionCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateDocumentationVersionCmd.MarkFlagRequired("documentation-version")
	apigateway_updateDocumentationVersionCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateDocumentationVersionCmd)
}
