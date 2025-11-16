package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateApiMappingCmd = &cobra.Command{
	Use:   "update-api-mapping",
	Short: "The API mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateApiMappingCmd).Standalone()

	apigatewayv2_updateApiMappingCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_updateApiMappingCmd.Flags().String("api-mapping-id", "", "The API mapping identifier.")
	apigatewayv2_updateApiMappingCmd.Flags().String("api-mapping-key", "", "The API mapping key.")
	apigatewayv2_updateApiMappingCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_updateApiMappingCmd.Flags().String("stage", "", "The API stage.")
	apigatewayv2_updateApiMappingCmd.MarkFlagRequired("api-id")
	apigatewayv2_updateApiMappingCmd.MarkFlagRequired("api-mapping-id")
	apigatewayv2_updateApiMappingCmd.MarkFlagRequired("domain-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateApiMappingCmd)
}
