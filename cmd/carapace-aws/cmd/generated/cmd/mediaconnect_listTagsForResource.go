package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags on a MediaConnect resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_listTagsForResourceCmd).Standalone()

		mediaconnect_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the MediaConnect resource for which to list the tags.")
		mediaconnect_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_listTagsForResourceCmd)
}
