package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getBasePathMappingsCmd = &cobra.Command{
	Use:   "get-base-path-mappings",
	Short: "Represents a collection of BasePathMapping resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getBasePathMappingsCmd).Standalone()

	apigateway_getBasePathMappingsCmd.Flags().String("domain-name", "", "The domain name of a BasePathMapping resource.")
	apigateway_getBasePathMappingsCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
	apigateway_getBasePathMappingsCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getBasePathMappingsCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigateway_getBasePathMappingsCmd.MarkFlagRequired("domain-name")
	apigatewayCmd.AddCommand(apigateway_getBasePathMappingsCmd)
}
