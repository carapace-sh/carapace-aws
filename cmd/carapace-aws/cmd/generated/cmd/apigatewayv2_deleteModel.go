package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteModelCmd = &cobra.Command{
	Use:   "delete-model",
	Short: "Deletes a Model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteModelCmd).Standalone()

	apigatewayv2_deleteModelCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteModelCmd.Flags().String("model-id", "", "The model ID.")
	apigatewayv2_deleteModelCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteModelCmd.MarkFlagRequired("model-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteModelCmd)
}
