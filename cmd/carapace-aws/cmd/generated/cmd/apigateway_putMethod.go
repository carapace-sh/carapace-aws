package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_putMethodCmd = &cobra.Command{
	Use:   "put-method",
	Short: "Add a method to an existing Resource resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_putMethodCmd).Standalone()

	apigateway_putMethodCmd.Flags().Bool("api-key-required", false, "Specifies whether the method required a valid ApiKey.")
	apigateway_putMethodCmd.Flags().String("authorization-scopes", "", "A list of authorization scopes configured on the method.")
	apigateway_putMethodCmd.Flags().String("authorization-type", "", "The method's authorization type.")
	apigateway_putMethodCmd.Flags().String("authorizer-id", "", "Specifies the identifier of an Authorizer to use on this Method, if the type is CUSTOM or COGNITO\\_USER\\_POOLS.")
	apigateway_putMethodCmd.Flags().String("http-method", "", "Specifies the method request's HTTP method type.")
	apigateway_putMethodCmd.Flags().Bool("no-api-key-required", false, "Specifies whether the method required a valid ApiKey.")
	apigateway_putMethodCmd.Flags().String("operation-name", "", "A human-friendly operation identifier for the method.")
	apigateway_putMethodCmd.Flags().String("request-models", "", "Specifies the Model resources used for the request's content type.")
	apigateway_putMethodCmd.Flags().String("request-parameters", "", "A key-value map defining required or optional method request parameters that can be accepted by API Gateway.")
	apigateway_putMethodCmd.Flags().String("request-validator-id", "", "The identifier of a RequestValidator for validating the method request.")
	apigateway_putMethodCmd.Flags().String("resource-id", "", "The Resource identifier for the new Method resource.")
	apigateway_putMethodCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_putMethodCmd.MarkFlagRequired("authorization-type")
	apigateway_putMethodCmd.MarkFlagRequired("http-method")
	apigateway_putMethodCmd.Flag("no-api-key-required").Hidden = true
	apigateway_putMethodCmd.MarkFlagRequired("resource-id")
	apigateway_putMethodCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_putMethodCmd)
}
