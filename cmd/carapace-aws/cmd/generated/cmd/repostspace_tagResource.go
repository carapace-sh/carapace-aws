package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates tags with an AWS re:Post Private resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_tagResourceCmd).Standalone()

		repostspace_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that the tag is associated with.")
		repostspace_tagResourceCmd.Flags().String("tags", "", "The list of tag keys and values that must be associated with the resource.")
		repostspace_tagResourceCmd.MarkFlagRequired("resource-arn")
		repostspace_tagResourceCmd.MarkFlagRequired("tags")
	})
	repostspaceCmd.AddCommand(repostspace_tagResourceCmd)
}
