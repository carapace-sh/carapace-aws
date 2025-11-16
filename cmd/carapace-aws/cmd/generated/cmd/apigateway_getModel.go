package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getModelCmd = &cobra.Command{
	Use:   "get-model",
	Short: "Describes an existing model defined for a RestApi resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getModelCmd).Standalone()

		apigateway_getModelCmd.Flags().Bool("flatten", false, "A query parameter of a Boolean value to resolve (`true`) all external model references and returns a flattened model schema or not (`false`) The default is `false`.")
		apigateway_getModelCmd.Flags().String("model-name", "", "The name of the model as an identifier.")
		apigateway_getModelCmd.Flags().Bool("no-flatten", false, "A query parameter of a Boolean value to resolve (`true`) all external model references and returns a flattened model schema or not (`false`) The default is `false`.")
		apigateway_getModelCmd.Flags().String("rest-api-id", "", "The RestApi identifier under which the Model exists.")
		apigateway_getModelCmd.MarkFlagRequired("model-name")
		apigateway_getModelCmd.Flag("no-flatten").Hidden = true
		apigateway_getModelCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getModelCmd)
}
