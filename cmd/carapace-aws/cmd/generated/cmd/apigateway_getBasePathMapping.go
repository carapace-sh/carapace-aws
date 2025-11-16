package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getBasePathMappingCmd = &cobra.Command{
	Use:   "get-base-path-mapping",
	Short: "Describe a BasePathMapping resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getBasePathMappingCmd).Standalone()

	apigateway_getBasePathMappingCmd.Flags().String("base-path", "", "The base path name that callers of the API must provide as part of the URL after the domain name.")
	apigateway_getBasePathMappingCmd.Flags().String("domain-name", "", "The domain name of the BasePathMapping resource to be described.")
	apigateway_getBasePathMappingCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
	apigateway_getBasePathMappingCmd.MarkFlagRequired("base-path")
	apigateway_getBasePathMappingCmd.MarkFlagRequired("domain-name")
	apigatewayCmd.AddCommand(apigateway_getBasePathMappingCmd)
}
