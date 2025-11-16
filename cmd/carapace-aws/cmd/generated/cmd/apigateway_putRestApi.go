package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_putRestApiCmd = &cobra.Command{
	Use:   "put-rest-api",
	Short: "A feature of the API Gateway control service for updating an existing API with an input of external API definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_putRestApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_putRestApiCmd).Standalone()

		apigateway_putRestApiCmd.Flags().String("body", "", "The PUT request body containing external API definitions.")
		apigateway_putRestApiCmd.Flags().Bool("fail-on-warnings", false, "A query parameter to indicate whether to rollback the API update (`true`) or not (`false`) when a warning is encountered.")
		apigateway_putRestApiCmd.Flags().String("mode", "", "The `mode` query parameter to specify the update mode.")
		apigateway_putRestApiCmd.Flags().Bool("no-fail-on-warnings", false, "A query parameter to indicate whether to rollback the API update (`true`) or not (`false`) when a warning is encountered.")
		apigateway_putRestApiCmd.Flags().String("parameters", "", "Custom header parameters as part of the request.")
		apigateway_putRestApiCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_putRestApiCmd.MarkFlagRequired("body")
		apigateway_putRestApiCmd.Flag("no-fail-on-warnings").Hidden = true
		apigateway_putRestApiCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_putRestApiCmd)
}
