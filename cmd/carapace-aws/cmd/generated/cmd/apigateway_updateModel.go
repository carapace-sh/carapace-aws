package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateModelCmd = &cobra.Command{
	Use:   "update-model",
	Short: "Changes information about a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateModelCmd).Standalone()

		apigateway_updateModelCmd.Flags().String("model-name", "", "The name of the model to update.")
		apigateway_updateModelCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateModelCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_updateModelCmd.MarkFlagRequired("model-name")
		apigateway_updateModelCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_updateModelCmd)
}
