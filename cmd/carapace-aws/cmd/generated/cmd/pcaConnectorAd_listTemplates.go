package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Lists the templates, if any, that are associated with a connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_listTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_listTemplatesCmd).Standalone()

		pcaConnectorAd_listTemplatesCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
		pcaConnectorAd_listTemplatesCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response on each page.")
		pcaConnectorAd_listTemplatesCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
		pcaConnectorAd_listTemplatesCmd.MarkFlagRequired("connector-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_listTemplatesCmd)
}
