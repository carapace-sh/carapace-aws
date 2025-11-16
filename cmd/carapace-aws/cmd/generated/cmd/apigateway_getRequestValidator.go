package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getRequestValidatorCmd = &cobra.Command{
	Use:   "get-request-validator",
	Short: "Gets a RequestValidator of a given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getRequestValidatorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getRequestValidatorCmd).Standalone()

		apigateway_getRequestValidatorCmd.Flags().String("request-validator-id", "", "The identifier of the RequestValidator to be retrieved.")
		apigateway_getRequestValidatorCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getRequestValidatorCmd.MarkFlagRequired("request-validator-id")
		apigateway_getRequestValidatorCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getRequestValidatorCmd)
}
