package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Creates a Model for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createModelCmd).Standalone()

	apigatewayv2_createModelCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_createModelCmd.Flags().String("content-type", "", "The content-type for the model, for example, \"application/json\".")
	apigatewayv2_createModelCmd.Flags().String("description", "", "The description of the model.")
	apigatewayv2_createModelCmd.Flags().String("name", "", "The name of the model.")
	apigatewayv2_createModelCmd.Flags().String("schema", "", "The schema for the model.")
	apigatewayv2_createModelCmd.MarkFlagRequired("api-id")
	apigatewayv2_createModelCmd.MarkFlagRequired("name")
	apigatewayv2_createModelCmd.MarkFlagRequired("schema")
	apigatewayv2Cmd.AddCommand(apigatewayv2_createModelCmd)
}
