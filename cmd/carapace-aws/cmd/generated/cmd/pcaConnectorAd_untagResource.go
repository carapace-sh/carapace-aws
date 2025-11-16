package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from your resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_untagResourceCmd).Standalone()

	pcaConnectorAd_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that was returned when you created the resource.")
	pcaConnectorAd_untagResourceCmd.Flags().String("tag-keys", "", "Specifies a list of tag keys that you want to remove from the specified resources.")
	pcaConnectorAd_untagResourceCmd.MarkFlagRequired("resource-arn")
	pcaConnectorAd_untagResourceCmd.MarkFlagRequired("tag-keys")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_untagResourceCmd)
}
