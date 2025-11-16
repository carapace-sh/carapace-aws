package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_testInvokeMethodCmd = &cobra.Command{
	Use:   "test-invoke-method",
	Short: "Simulate the invocation of a Method in your RestApi with headers, parameters, and an incoming request body.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_testInvokeMethodCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_testInvokeMethodCmd).Standalone()

		apigateway_testInvokeMethodCmd.Flags().String("body", "", "The simulated request body of an incoming invocation request.")
		apigateway_testInvokeMethodCmd.Flags().String("client-certificate-id", "", "A ClientCertificate identifier to use in the test invocation.")
		apigateway_testInvokeMethodCmd.Flags().String("headers", "", "A key-value map of headers to simulate an incoming invocation request.")
		apigateway_testInvokeMethodCmd.Flags().String("http-method", "", "Specifies a test invoke method request's HTTP method.")
		apigateway_testInvokeMethodCmd.Flags().String("multi-value-headers", "", "The headers as a map from string to list of values to simulate an incoming invocation request.")
		apigateway_testInvokeMethodCmd.Flags().String("path-with-query-string", "", "The URI path, including query string, of the simulated invocation request.")
		apigateway_testInvokeMethodCmd.Flags().String("resource-id", "", "Specifies a test invoke method request's resource ID.")
		apigateway_testInvokeMethodCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_testInvokeMethodCmd.Flags().String("stage-variables", "", "A key-value map of stage variables to simulate an invocation on a deployed Stage.")
		apigateway_testInvokeMethodCmd.MarkFlagRequired("http-method")
		apigateway_testInvokeMethodCmd.MarkFlagRequired("resource-id")
		apigateway_testInvokeMethodCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_testInvokeMethodCmd)
}
