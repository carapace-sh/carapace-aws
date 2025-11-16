package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_exportApiCmd = &cobra.Command{
	Use:   "export-api",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_exportApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_exportApiCmd).Standalone()

		apigatewayv2_exportApiCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_exportApiCmd.Flags().String("export-version", "", "The version of the API Gateway export algorithm.")
		apigatewayv2_exportApiCmd.Flags().String("include-extensions", "", "Specifies whether to include [API Gateway extensions](https://docs.aws.amazon.com//apigateway/latest/developerguide/api-gateway-swagger-extensions.html) in the exported API definition.")
		apigatewayv2_exportApiCmd.Flags().String("output-type", "", "The output type of the exported definition file.")
		apigatewayv2_exportApiCmd.Flags().String("specification", "", "The version of the API specification to use.")
		apigatewayv2_exportApiCmd.Flags().String("stage-name", "", "The name of the API stage to export.")
		apigatewayv2_exportApiCmd.MarkFlagRequired("api-id")
		apigatewayv2_exportApiCmd.MarkFlagRequired("output-type")
		apigatewayv2_exportApiCmd.MarkFlagRequired("specification")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_exportApiCmd)
}
