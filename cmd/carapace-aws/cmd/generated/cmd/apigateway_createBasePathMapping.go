package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createBasePathMappingCmd = &cobra.Command{
	Use:   "create-base-path-mapping",
	Short: "Creates a new BasePathMapping resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createBasePathMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createBasePathMappingCmd).Standalone()

		apigateway_createBasePathMappingCmd.Flags().String("base-path", "", "The base path name that callers of the API must provide as part of the URL after the domain name.")
		apigateway_createBasePathMappingCmd.Flags().String("domain-name", "", "The domain name of the BasePathMapping resource to create.")
		apigateway_createBasePathMappingCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
		apigateway_createBasePathMappingCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_createBasePathMappingCmd.Flags().String("stage", "", "The name of the API's stage that you want to use for this mapping.")
		apigateway_createBasePathMappingCmd.MarkFlagRequired("domain-name")
		apigateway_createBasePathMappingCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_createBasePathMappingCmd)
}
