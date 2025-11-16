package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createRestApiCmd = &cobra.Command{
	Use:   "create-rest-api",
	Short: "Creates a new RestApi resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createRestApiCmd).Standalone()

	apigateway_createRestApiCmd.Flags().String("api-key-source", "", "The source of the API key for metering requests according to a usage plan.")
	apigateway_createRestApiCmd.Flags().String("binary-media-types", "", "The list of binary media types supported by the RestApi.")
	apigateway_createRestApiCmd.Flags().String("clone-from", "", "The ID of the RestApi that you want to clone from.")
	apigateway_createRestApiCmd.Flags().String("description", "", "The description of the RestApi.")
	apigateway_createRestApiCmd.Flags().Bool("disable-execute-api-endpoint", false, "Specifies whether clients can invoke your API by using the default `execute-api` endpoint.")
	apigateway_createRestApiCmd.Flags().String("endpoint-configuration", "", "The endpoint configuration of this RestApi showing the endpoint types and IP address types of the API.")
	apigateway_createRestApiCmd.Flags().String("minimum-compression-size", "", "A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API.")
	apigateway_createRestApiCmd.Flags().String("name", "", "The name of the RestApi.")
	apigateway_createRestApiCmd.Flags().Bool("no-disable-execute-api-endpoint", false, "Specifies whether clients can invoke your API by using the default `execute-api` endpoint.")
	apigateway_createRestApiCmd.Flags().String("policy", "", "A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration.")
	apigateway_createRestApiCmd.Flags().String("tags", "", "The key-value map of strings.")
	apigateway_createRestApiCmd.Flags().String("version", "", "A version identifier for the API.")
	apigateway_createRestApiCmd.MarkFlagRequired("name")
	apigateway_createRestApiCmd.Flag("no-disable-execute-api-endpoint").Hidden = true
	apigatewayCmd.AddCommand(apigateway_createRestApiCmd)
}
