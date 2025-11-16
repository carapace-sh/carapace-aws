package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getModelsCmd = &cobra.Command{
	Use:   "get-models",
	Short: "Describes existing Models defined for a RestApi resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getModelsCmd).Standalone()

		apigateway_getModelsCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getModelsCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getModelsCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getModelsCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getModelsCmd)
}
