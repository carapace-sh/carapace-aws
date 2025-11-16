package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_tagResourceCmd).Standalone()

		cloud9_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Cloud9 development environment to add tags to.")
		cloud9_tagResourceCmd.Flags().String("tags", "", "The list of tags to add to the given Cloud9 development environment.")
		cloud9_tagResourceCmd.MarkFlagRequired("resource-arn")
		cloud9_tagResourceCmd.MarkFlagRequired("tags")
	})
	cloud9Cmd.AddCommand(cloud9_tagResourceCmd)
}
