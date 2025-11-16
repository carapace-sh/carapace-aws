package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createDocumentationVersionCmd = &cobra.Command{
	Use:   "create-documentation-version",
	Short: "Creates a documentation version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createDocumentationVersionCmd).Standalone()

	apigateway_createDocumentationVersionCmd.Flags().String("description", "", "A description about the new documentation snapshot.")
	apigateway_createDocumentationVersionCmd.Flags().String("documentation-version", "", "The version identifier of the new snapshot.")
	apigateway_createDocumentationVersionCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_createDocumentationVersionCmd.Flags().String("stage-name", "", "The stage name to be associated with the new documentation snapshot.")
	apigateway_createDocumentationVersionCmd.MarkFlagRequired("documentation-version")
	apigateway_createDocumentationVersionCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_createDocumentationVersionCmd)
}
