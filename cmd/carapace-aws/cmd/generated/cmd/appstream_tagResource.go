package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites one or more tags for the specified AppStream 2.0 resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_tagResourceCmd).Standalone()

		appstream_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		appstream_tagResourceCmd.Flags().String("tags", "", "The tags to associate.")
		appstream_tagResourceCmd.MarkFlagRequired("resource-arn")
		appstream_tagResourceCmd.MarkFlagRequired("tags")
	})
	appstreamCmd.AddCommand(appstream_tagResourceCmd)
}
