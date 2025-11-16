package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets a list of the tags associated with an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_listTagsForResourceCmd).Standalone()

		cloud9_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Cloud9 development environment to get the tags for.")
		cloud9_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	cloud9Cmd.AddCommand(cloud9_listTagsForResourceCmd)
}
