package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets information about AWS tags for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_listTagsForResourceCmd).Standalone()

		ivsRealtime_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be retrieved.")
		ivsRealtime_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_listTagsForResourceCmd)
}
