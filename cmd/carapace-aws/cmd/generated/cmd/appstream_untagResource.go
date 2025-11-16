package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Disassociates one or more specified tags from the specified AppStream 2.0 resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_untagResourceCmd).Standalone()

		appstream_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		appstream_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys for the tags to disassociate.")
		appstream_untagResourceCmd.MarkFlagRequired("resource-arn")
		appstream_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	appstreamCmd.AddCommand(appstream_untagResourceCmd)
}
