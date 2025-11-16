package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Adds a new Model resource to an existing RestApi resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createModelCmd).Standalone()

		apigateway_createModelCmd.Flags().String("content-type", "", "The content-type for the model.")
		apigateway_createModelCmd.Flags().String("description", "", "The description of the model.")
		apigateway_createModelCmd.Flags().String("name", "", "The name of the model.")
		apigateway_createModelCmd.Flags().String("rest-api-id", "", "The RestApi identifier under which the Model will be created.")
		apigateway_createModelCmd.Flags().String("schema", "", "The schema for the model.")
		apigateway_createModelCmd.MarkFlagRequired("content-type")
		apigateway_createModelCmd.MarkFlagRequired("name")
		apigateway_createModelCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_createModelCmd)
}
