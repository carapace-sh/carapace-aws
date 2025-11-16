package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_testInvokeAuthorizerCmd = &cobra.Command{
	Use:   "test-invoke-authorizer",
	Short: "Tests a custom authorization behavior by invoking a specified custom authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_testInvokeAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_testInvokeAuthorizerCmd).Standalone()

		iot_testInvokeAuthorizerCmd.Flags().String("authorizer-name", "", "The custom authorizer name.")
		iot_testInvokeAuthorizerCmd.Flags().String("http-context", "", "Specifies a test HTTP authorization request.")
		iot_testInvokeAuthorizerCmd.Flags().String("mqtt-context", "", "Specifies a test MQTT authorization request.")
		iot_testInvokeAuthorizerCmd.Flags().String("tls-context", "", "Specifies a test TLS authorization request.")
		iot_testInvokeAuthorizerCmd.Flags().String("token", "", "The token returned by your custom authentication service.")
		iot_testInvokeAuthorizerCmd.Flags().String("token-signature", "", "The signature made with the token and your custom authentication service's private key.")
		iot_testInvokeAuthorizerCmd.MarkFlagRequired("authorizer-name")
	})
	iotCmd.AddCommand(iot_testInvokeAuthorizerCmd)
}
