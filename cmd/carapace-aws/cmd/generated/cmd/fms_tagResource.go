package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_tagResourceCmd).Standalone()

		fms_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to return tags for.")
		fms_tagResourceCmd.Flags().String("tag-list", "", "The tags to add to the resource.")
		fms_tagResourceCmd.MarkFlagRequired("resource-arn")
		fms_tagResourceCmd.MarkFlagRequired("tag-list")
	})
	fmsCmd.AddCommand(fms_tagResourceCmd)
}
