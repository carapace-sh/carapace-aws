package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_importDocumentationPartsCmd = &cobra.Command{
	Use:   "import-documentation-parts",
	Short: "Imports documentation parts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_importDocumentationPartsCmd).Standalone()

	apigateway_importDocumentationPartsCmd.Flags().String("body", "", "Raw byte array representing the to-be-imported documentation parts.")
	apigateway_importDocumentationPartsCmd.Flags().Bool("fail-on-warnings", false, "A query parameter to specify whether to rollback the documentation importation (`true`) or not (`false`) when a warning is encountered.")
	apigateway_importDocumentationPartsCmd.Flags().String("mode", "", "A query parameter to indicate whether to overwrite (`overwrite`) any existing DocumentationParts definition or to merge (`merge`) the new definition into the existing one.")
	apigateway_importDocumentationPartsCmd.Flags().Bool("no-fail-on-warnings", false, "A query parameter to specify whether to rollback the documentation importation (`true`) or not (`false`) when a warning is encountered.")
	apigateway_importDocumentationPartsCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_importDocumentationPartsCmd.MarkFlagRequired("body")
	apigateway_importDocumentationPartsCmd.Flag("no-fail-on-warnings").Hidden = true
	apigateway_importDocumentationPartsCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_importDocumentationPartsCmd)
}
