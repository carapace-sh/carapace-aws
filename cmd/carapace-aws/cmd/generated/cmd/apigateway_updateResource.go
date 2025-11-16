package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateResourceCmd = &cobra.Command{
	Use:   "update-resource",
	Short: "Changes information about a Resource resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateResourceCmd).Standalone()

	apigateway_updateResourceCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateResourceCmd.Flags().String("resource-id", "", "The identifier of the Resource resource.")
	apigateway_updateResourceCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateResourceCmd.MarkFlagRequired("resource-id")
	apigateway_updateResourceCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateResourceCmd)
}
