package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateApiKeyCmd = &cobra.Command{
	Use:   "update-api-key",
	Short: "Changes information about an ApiKey resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateApiKeyCmd).Standalone()

	apigateway_updateApiKeyCmd.Flags().String("api-key", "", "The identifier of the ApiKey resource to be updated.")
	apigateway_updateApiKeyCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateApiKeyCmd.MarkFlagRequired("api-key")
	apigatewayCmd.AddCommand(apigateway_updateApiKeyCmd)
}
