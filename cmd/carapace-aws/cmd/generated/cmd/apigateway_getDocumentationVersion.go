package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDocumentationVersionCmd = &cobra.Command{
	Use:   "get-documentation-version",
	Short: "Gets a documentation version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDocumentationVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getDocumentationVersionCmd).Standalone()

		apigateway_getDocumentationVersionCmd.Flags().String("documentation-version", "", "The version identifier of the to-be-retrieved documentation snapshot.")
		apigateway_getDocumentationVersionCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getDocumentationVersionCmd.MarkFlagRequired("documentation-version")
		apigateway_getDocumentationVersionCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getDocumentationVersionCmd)
}
