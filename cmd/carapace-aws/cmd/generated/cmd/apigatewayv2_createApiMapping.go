package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createApiMappingCmd = &cobra.Command{
	Use:   "create-api-mapping",
	Short: "Creates an API mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createApiMappingCmd).Standalone()

	apigatewayv2_createApiMappingCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_createApiMappingCmd.Flags().String("api-mapping-key", "", "The API mapping key.")
	apigatewayv2_createApiMappingCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_createApiMappingCmd.Flags().String("stage", "", "The API stage.")
	apigatewayv2_createApiMappingCmd.MarkFlagRequired("api-id")
	apigatewayv2_createApiMappingCmd.MarkFlagRequired("domain-name")
	apigatewayv2_createApiMappingCmd.MarkFlagRequired("stage")
	apigatewayv2Cmd.AddCommand(apigatewayv2_createApiMappingCmd)
}
