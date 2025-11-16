package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getModelCmd = &cobra.Command{
	Use:   "get-model",
	Short: "Gets a Model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getModelCmd).Standalone()

	apigatewayv2_getModelCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getModelCmd.Flags().String("model-id", "", "The model ID.")
	apigatewayv2_getModelCmd.MarkFlagRequired("api-id")
	apigatewayv2_getModelCmd.MarkFlagRequired("model-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getModelCmd)
}
