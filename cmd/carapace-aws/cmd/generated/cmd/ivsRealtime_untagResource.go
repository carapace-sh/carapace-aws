package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_untagResourceCmd).Standalone()

	ivsRealtime_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be untagged.")
	ivsRealtime_untagResourceCmd.Flags().String("tag-keys", "", "Array of tag keys (strings) for the tags to be removed.")
	ivsRealtime_untagResourceCmd.MarkFlagRequired("resource-arn")
	ivsRealtime_untagResourceCmd.MarkFlagRequired("tag-keys")
	ivsRealtimeCmd.AddCommand(ivsRealtime_untagResourceCmd)
}
