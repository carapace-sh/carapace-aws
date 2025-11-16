package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "A list of tags that are associated with this resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_listTagsForResourceCmd).Standalone()

		mediatailor_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with this resource.")
		mediatailor_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mediatailorCmd.AddCommand(mediatailor_listTagsForResourceCmd)
}
