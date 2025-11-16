package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "The resource to tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_tagResourceCmd).Standalone()

		mediatailor_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource.")
		mediatailor_tagResourceCmd.Flags().String("tags", "", "The tags to assign to the resource.")
		mediatailor_tagResourceCmd.MarkFlagRequired("resource-arn")
		mediatailor_tagResourceCmd.MarkFlagRequired("tags")
	})
	mediatailorCmd.AddCommand(mediatailor_tagResourceCmd)
}
