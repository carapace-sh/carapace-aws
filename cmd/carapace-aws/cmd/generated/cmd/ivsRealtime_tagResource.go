package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for the AWS resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_tagResourceCmd).Standalone()

	ivsRealtime_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be tagged.")
	ivsRealtime_tagResourceCmd.Flags().String("tags", "", "Array of tags to be added or updated.")
	ivsRealtime_tagResourceCmd.MarkFlagRequired("resource-arn")
	ivsRealtime_tagResourceCmd.MarkFlagRequired("tags")
	ivsRealtimeCmd.AddCommand(ivsRealtime_tagResourceCmd)
}
