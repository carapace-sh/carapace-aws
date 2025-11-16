package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_untagResourceCmd).Standalone()

	datasync_untagResourceCmd.Flags().String("keys", "", "Specifies the keys in the tags that you want to remove.")
	datasync_untagResourceCmd.Flags().String("resource-arn", "", "Specifies the Amazon Resource Name (ARN) of the resource to remove the tags from.")
	datasync_untagResourceCmd.MarkFlagRequired("keys")
	datasync_untagResourceCmd.MarkFlagRequired("resource-arn")
	datasyncCmd.AddCommand(datasync_untagResourceCmd)
}
