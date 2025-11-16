package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to your resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_tagResourceCmd).Standalone()

		pcaConnectorAd_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that was returned when you created the resource.")
		pcaConnectorAd_tagResourceCmd.Flags().String("tags", "", "Metadata assigned to a directory registration consisting of a key-value pair.")
		pcaConnectorAd_tagResourceCmd.MarkFlagRequired("resource-arn")
		pcaConnectorAd_tagResourceCmd.MarkFlagRequired("tags")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_tagResourceCmd)
}
