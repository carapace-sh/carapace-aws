package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createResourceCmd = &cobra.Command{
	Use:   "create-resource",
	Short: "Creates a Resource resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createResourceCmd).Standalone()

		apigateway_createResourceCmd.Flags().String("parent-id", "", "The parent resource's identifier.")
		apigateway_createResourceCmd.Flags().String("path-part", "", "The last path segment for this resource.")
		apigateway_createResourceCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_createResourceCmd.MarkFlagRequired("parent-id")
		apigateway_createResourceCmd.MarkFlagRequired("path-part")
		apigateway_createResourceCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_createResourceCmd)
}
