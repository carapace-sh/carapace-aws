package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Lists the connectors that you created by using the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API\\_CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector) action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_listConnectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_listConnectorsCmd).Standalone()

		pcaConnectorAd_listConnectorsCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response on each page.")
		pcaConnectorAd_listConnectorsCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_listConnectorsCmd)
}
