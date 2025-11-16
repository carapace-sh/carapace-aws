package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateDocumentationPartCmd = &cobra.Command{
	Use:   "update-documentation-part",
	Short: "Updates a documentation part.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateDocumentationPartCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateDocumentationPartCmd).Standalone()

		apigateway_updateDocumentationPartCmd.Flags().String("documentation-part-id", "", "The identifier of the to-be-updated documentation part.")
		apigateway_updateDocumentationPartCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateDocumentationPartCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_updateDocumentationPartCmd.MarkFlagRequired("documentation-part-id")
		apigateway_updateDocumentationPartCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_updateDocumentationPartCmd)
}
