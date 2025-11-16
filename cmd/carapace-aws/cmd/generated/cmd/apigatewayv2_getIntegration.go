package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getIntegrationCmd = &cobra.Command{
	Use:   "get-integration",
	Short: "Gets an Integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getIntegrationCmd).Standalone()

	apigatewayv2_getIntegrationCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getIntegrationCmd.Flags().String("integration-id", "", "The integration ID.")
	apigatewayv2_getIntegrationCmd.MarkFlagRequired("api-id")
	apigatewayv2_getIntegrationCmd.MarkFlagRequired("integration-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getIntegrationCmd)
}
