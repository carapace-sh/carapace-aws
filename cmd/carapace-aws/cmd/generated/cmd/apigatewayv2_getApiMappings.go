package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getApiMappingsCmd = &cobra.Command{
	Use:   "get-api-mappings",
	Short: "Gets API mappings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getApiMappingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getApiMappingsCmd).Standalone()

		apigatewayv2_getApiMappingsCmd.Flags().String("domain-name", "", "The domain name.")
		apigatewayv2_getApiMappingsCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getApiMappingsCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getApiMappingsCmd.MarkFlagRequired("domain-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getApiMappingsCmd)
}
