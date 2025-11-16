package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd = &cobra.Command{
	Use:   "list-template-group-access-control-entries",
	Short: "Lists group access control entries you created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd).Standalone()

		pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response on each page.")
		pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
		pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
		pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd.MarkFlagRequired("template-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_listTemplateGroupAccessControlEntriesCmd)
}
