package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getModelTemplateCmd = &cobra.Command{
	Use:   "get-model-template",
	Short: "Generates a sample mapping template that can be used to transform a payload into the structure of a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getModelTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getModelTemplateCmd).Standalone()

		apigateway_getModelTemplateCmd.Flags().String("model-name", "", "The name of the model for which to generate a template.")
		apigateway_getModelTemplateCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getModelTemplateCmd.MarkFlagRequired("model-name")
		apigateway_getModelTemplateCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getModelTemplateCmd)
}
