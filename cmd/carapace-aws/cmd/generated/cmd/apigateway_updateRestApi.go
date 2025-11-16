package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateRestApiCmd = &cobra.Command{
	Use:   "update-rest-api",
	Short: "Changes information about the specified API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateRestApiCmd).Standalone()

	apigateway_updateRestApiCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateRestApiCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateRestApiCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateRestApiCmd)
}
