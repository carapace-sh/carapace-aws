package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getApiMappingCmd = &cobra.Command{
	Use:   "get-api-mapping",
	Short: "Gets an API mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getApiMappingCmd).Standalone()

	apigatewayv2_getApiMappingCmd.Flags().String("api-mapping-id", "", "The API mapping identifier.")
	apigatewayv2_getApiMappingCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_getApiMappingCmd.MarkFlagRequired("api-mapping-id")
	apigatewayv2_getApiMappingCmd.MarkFlagRequired("domain-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getApiMappingCmd)
}
