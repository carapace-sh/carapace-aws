package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteModelCmd = &cobra.Command{
	Use:   "delete-model",
	Short: "Deletes a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteModelCmd).Standalone()

	apigateway_deleteModelCmd.Flags().String("model-name", "", "The name of the model to delete.")
	apigateway_deleteModelCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_deleteModelCmd.MarkFlagRequired("model-name")
	apigateway_deleteModelCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_deleteModelCmd)
}
