package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getRequestValidatorsCmd = &cobra.Command{
	Use:   "get-request-validators",
	Short: "Gets the RequestValidators collection of a given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getRequestValidatorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getRequestValidatorsCmd).Standalone()

		apigateway_getRequestValidatorsCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getRequestValidatorsCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getRequestValidatorsCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getRequestValidatorsCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getRequestValidatorsCmd)
}
