package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from an Amazon Q Business application or a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_untagResourceCmd).Standalone()

	qbusiness_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Q Business application, or data source to remove the tag from.")
	qbusiness_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the Amazon Q Business application or data source.")
	qbusiness_untagResourceCmd.MarkFlagRequired("resource-arn")
	qbusiness_untagResourceCmd.MarkFlagRequired("tag-keys")
	qbusinessCmd.AddCommand(qbusiness_untagResourceCmd)
}
