package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_importRestApiCmd = &cobra.Command{
	Use:   "import-rest-api",
	Short: "A feature of the API Gateway control service for creating a new API from an external API definition file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_importRestApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_importRestApiCmd).Standalone()

		apigateway_importRestApiCmd.Flags().String("body", "", "The POST request body containing external API definitions.")
		apigateway_importRestApiCmd.Flags().Bool("fail-on-warnings", false, "A query parameter to indicate whether to rollback the API creation (`true`) or not (`false`) when a warning is encountered.")
		apigateway_importRestApiCmd.Flags().Bool("no-fail-on-warnings", false, "A query parameter to indicate whether to rollback the API creation (`true`) or not (`false`) when a warning is encountered.")
		apigateway_importRestApiCmd.Flags().String("parameters", "", "A key-value map of context-specific query string parameters specifying the behavior of different API importing operations.")
		apigateway_importRestApiCmd.MarkFlagRequired("body")
		apigateway_importRestApiCmd.Flag("no-fail-on-warnings").Hidden = true
	})
	apigatewayCmd.AddCommand(apigateway_importRestApiCmd)
}
