package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getSdkCmd = &cobra.Command{
	Use:   "get-sdk",
	Short: "Generates a client SDK for a RestApi and Stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getSdkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getSdkCmd).Standalone()

		apigateway_getSdkCmd.Flags().String("parameters", "", "A string-to-string key-value map of query parameters `sdkType`-dependent properties of the SDK.")
		apigateway_getSdkCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getSdkCmd.Flags().String("sdk-type", "", "The language for the generated SDK.")
		apigateway_getSdkCmd.Flags().String("stage-name", "", "The name of the Stage that the SDK will use.")
		apigateway_getSdkCmd.MarkFlagRequired("rest-api-id")
		apigateway_getSdkCmd.MarkFlagRequired("sdk-type")
		apigateway_getSdkCmd.MarkFlagRequired("stage-name")
	})
	apigatewayCmd.AddCommand(apigateway_getSdkCmd)
}
