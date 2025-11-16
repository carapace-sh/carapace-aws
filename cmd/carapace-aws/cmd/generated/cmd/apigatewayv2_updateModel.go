package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateModelCmd = &cobra.Command{
	Use:   "update-model",
	Short: "Updates a Model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateModelCmd).Standalone()

	apigatewayv2_updateModelCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_updateModelCmd.Flags().String("content-type", "", "The content-type for the model, for example, \"application/json\".")
	apigatewayv2_updateModelCmd.Flags().String("description", "", "The description of the model.")
	apigatewayv2_updateModelCmd.Flags().String("model-id", "", "The model ID.")
	apigatewayv2_updateModelCmd.Flags().String("name", "", "The name of the model.")
	apigatewayv2_updateModelCmd.Flags().String("schema", "", "The schema for the model.")
	apigatewayv2_updateModelCmd.MarkFlagRequired("api-id")
	apigatewayv2_updateModelCmd.MarkFlagRequired("model-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateModelCmd)
}
