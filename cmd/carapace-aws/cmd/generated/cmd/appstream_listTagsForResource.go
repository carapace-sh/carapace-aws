package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves a list of all tags for the specified AppStream 2.0 resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_listTagsForResourceCmd).Standalone()

		appstream_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		appstream_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	appstreamCmd.AddCommand(appstream_listTagsForResourceCmd)
}
