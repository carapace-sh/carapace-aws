package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateBasePathMappingCmd = &cobra.Command{
	Use:   "update-base-path-mapping",
	Short: "Changes information about the BasePathMapping resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateBasePathMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateBasePathMappingCmd).Standalone()

		apigateway_updateBasePathMappingCmd.Flags().String("base-path", "", "The base path of the BasePathMapping resource to change.")
		apigateway_updateBasePathMappingCmd.Flags().String("domain-name", "", "The domain name of the BasePathMapping resource to change.")
		apigateway_updateBasePathMappingCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
		apigateway_updateBasePathMappingCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateBasePathMappingCmd.MarkFlagRequired("base-path")
		apigateway_updateBasePathMappingCmd.MarkFlagRequired("domain-name")
	})
	apigatewayCmd.AddCommand(apigateway_updateBasePathMappingCmd)
}
