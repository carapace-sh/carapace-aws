package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags, if any, that are associated with your resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_listTagsForResourceCmd).Standalone()

		pcaConnectorAd_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that was returned when you created the resource.")
		pcaConnectorAd_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_listTagsForResourceCmd)
}
