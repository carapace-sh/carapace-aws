package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_importApiCmd = &cobra.Command{
	Use:   "import-api",
	Short: "Imports an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_importApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_importApiCmd).Standalone()

		apigatewayv2_importApiCmd.Flags().String("basepath", "", "Specifies how to interpret the base path of the API during import.")
		apigatewayv2_importApiCmd.Flags().String("body", "", "The OpenAPI definition.")
		apigatewayv2_importApiCmd.Flags().String("fail-on-warnings", "", "Specifies whether to rollback the API creation when a warning is encountered.")
		apigatewayv2_importApiCmd.MarkFlagRequired("body")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_importApiCmd)
}
