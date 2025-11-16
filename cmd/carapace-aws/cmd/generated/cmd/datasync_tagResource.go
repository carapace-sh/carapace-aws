package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies a *tag* to an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_tagResourceCmd).Standalone()

		datasync_tagResourceCmd.Flags().String("resource-arn", "", "Specifies the Amazon Resource Name (ARN) of the resource to apply the tag to.")
		datasync_tagResourceCmd.Flags().String("tags", "", "Specifies the tags that you want to apply to the resource.")
		datasync_tagResourceCmd.MarkFlagRequired("resource-arn")
		datasync_tagResourceCmd.MarkFlagRequired("tags")
	})
	datasyncCmd.AddCommand(datasync_tagResourceCmd)
}
