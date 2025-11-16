package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from your resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorScep_untagResourceCmd).Standalone()

		pcaConnectorScep_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		pcaConnectorScep_untagResourceCmd.Flags().String("tag-keys", "", "Specifies a list of tag keys that you want to remove from the specified resources.")
		pcaConnectorScep_untagResourceCmd.MarkFlagRequired("resource-arn")
		pcaConnectorScep_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_untagResourceCmd)
}
