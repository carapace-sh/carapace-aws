package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getModelsCmd = &cobra.Command{
	Use:   "get-models",
	Short: "Gets the Models for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getModelsCmd).Standalone()

	apigatewayv2_getModelsCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getModelsCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
	apigatewayv2_getModelsCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
	apigatewayv2_getModelsCmd.MarkFlagRequired("api-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getModelsCmd)
}
