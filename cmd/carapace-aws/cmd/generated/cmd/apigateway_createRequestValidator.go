package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createRequestValidatorCmd = &cobra.Command{
	Use:   "create-request-validator",
	Short: "Creates a RequestValidator of a given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createRequestValidatorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createRequestValidatorCmd).Standalone()

		apigateway_createRequestValidatorCmd.Flags().String("name", "", "The name of the to-be-created RequestValidator.")
		apigateway_createRequestValidatorCmd.Flags().Bool("no-validate-request-body", false, "A Boolean flag to indicate whether to validate request body according to the configured model schema for the method (`true`) or not (`false`).")
		apigateway_createRequestValidatorCmd.Flags().Bool("no-validate-request-parameters", false, "A Boolean flag to indicate whether to validate request parameters, `true`, or not `false`.")
		apigateway_createRequestValidatorCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_createRequestValidatorCmd.Flags().Bool("validate-request-body", false, "A Boolean flag to indicate whether to validate request body according to the configured model schema for the method (`true`) or not (`false`).")
		apigateway_createRequestValidatorCmd.Flags().Bool("validate-request-parameters", false, "A Boolean flag to indicate whether to validate request parameters, `true`, or not `false`.")
		apigateway_createRequestValidatorCmd.Flag("no-validate-request-body").Hidden = true
		apigateway_createRequestValidatorCmd.Flag("no-validate-request-parameters").Hidden = true
		apigateway_createRequestValidatorCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_createRequestValidatorCmd)
}
