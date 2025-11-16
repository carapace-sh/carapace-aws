package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getModelTemplateCmd = &cobra.Command{
	Use:   "get-model-template",
	Short: "Gets a model template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getModelTemplateCmd).Standalone()

	apigatewayv2_getModelTemplateCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getModelTemplateCmd.Flags().String("model-id", "", "The model ID.")
	apigatewayv2_getModelTemplateCmd.MarkFlagRequired("api-id")
	apigatewayv2_getModelTemplateCmd.MarkFlagRequired("model-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getModelTemplateCmd)
}
