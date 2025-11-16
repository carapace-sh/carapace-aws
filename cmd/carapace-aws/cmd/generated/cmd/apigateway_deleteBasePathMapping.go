package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteBasePathMappingCmd = &cobra.Command{
	Use:   "delete-base-path-mapping",
	Short: "Deletes the BasePathMapping resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteBasePathMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteBasePathMappingCmd).Standalone()

		apigateway_deleteBasePathMappingCmd.Flags().String("base-path", "", "The base path name of the BasePathMapping resource to delete.")
		apigateway_deleteBasePathMappingCmd.Flags().String("domain-name", "", "The domain name of the BasePathMapping resource to delete.")
		apigateway_deleteBasePathMappingCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
		apigateway_deleteBasePathMappingCmd.MarkFlagRequired("base-path")
		apigateway_deleteBasePathMappingCmd.MarkFlagRequired("domain-name")
	})
	apigatewayCmd.AddCommand(apigateway_deleteBasePathMappingCmd)
}
