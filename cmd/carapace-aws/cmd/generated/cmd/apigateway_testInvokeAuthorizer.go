package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_testInvokeAuthorizerCmd = &cobra.Command{
	Use:   "test-invoke-authorizer",
	Short: "Simulate the execution of an Authorizer in your RestApi with headers, parameters, and an incoming request body.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_testInvokeAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_testInvokeAuthorizerCmd).Standalone()

		apigateway_testInvokeAuthorizerCmd.Flags().String("additional-context", "", "A key-value map of additional context variables.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("authorizer-id", "", "Specifies a test invoke authorizer request's Authorizer ID.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("body", "", "The simulated request body of an incoming invocation request.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("headers", "", "A key-value map of headers to simulate an incoming invocation request.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("multi-value-headers", "", "The headers as a map from string to list of values to simulate an incoming invocation request.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("path-with-query-string", "", "The URI path, including query string, of the simulated invocation request.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_testInvokeAuthorizerCmd.Flags().String("stage-variables", "", "A key-value map of stage variables to simulate an invocation on a deployed Stage.")
		apigateway_testInvokeAuthorizerCmd.MarkFlagRequired("authorizer-id")
		apigateway_testInvokeAuthorizerCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_testInvokeAuthorizerCmd)
}
