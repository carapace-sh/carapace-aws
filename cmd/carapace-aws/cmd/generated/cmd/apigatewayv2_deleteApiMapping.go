package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteApiMappingCmd = &cobra.Command{
	Use:   "delete-api-mapping",
	Short: "Deletes an API mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteApiMappingCmd).Standalone()

	apigatewayv2_deleteApiMappingCmd.Flags().String("api-mapping-id", "", "The API mapping identifier.")
	apigatewayv2_deleteApiMappingCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_deleteApiMappingCmd.MarkFlagRequired("api-mapping-id")
	apigatewayv2_deleteApiMappingCmd.MarkFlagRequired("domain-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteApiMappingCmd)
}
