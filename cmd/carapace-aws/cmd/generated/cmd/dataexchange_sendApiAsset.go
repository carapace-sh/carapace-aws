package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_sendApiAssetCmd = &cobra.Command{
	Use:   "send-api-asset",
	Short: "This operation invokes an API Gateway API asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_sendApiAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_sendApiAssetCmd).Standalone()

		dataexchange_sendApiAssetCmd.Flags().String("asset-id", "", "Asset ID value for the API request.")
		dataexchange_sendApiAssetCmd.Flags().String("body", "", "The request body.")
		dataexchange_sendApiAssetCmd.Flags().String("data-set-id", "", "Data set ID value for the API request.")
		dataexchange_sendApiAssetCmd.Flags().String("method", "", "HTTP method value for the API request.")
		dataexchange_sendApiAssetCmd.Flags().String("path", "", "URI path value for the API request.")
		dataexchange_sendApiAssetCmd.Flags().String("query-string-parameters", "", "Attach query string parameters to the end of the URI (for example, /v1/examplePath?exampleParam=exampleValue).")
		dataexchange_sendApiAssetCmd.Flags().String("request-headers", "", "Any header value prefixed with x-amzn-dataexchange-header- will have that stripped before sending the Asset API request.")
		dataexchange_sendApiAssetCmd.Flags().String("revision-id", "", "Revision ID value for the API request.")
		dataexchange_sendApiAssetCmd.MarkFlagRequired("asset-id")
		dataexchange_sendApiAssetCmd.MarkFlagRequired("data-set-id")
		dataexchange_sendApiAssetCmd.MarkFlagRequired("revision-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_sendApiAssetCmd)
}
