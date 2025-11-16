package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pipes_listTagsForResourceCmd).Standalone()

		pipes_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the pipe for which you want to view tags.")
		pipes_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	pipesCmd.AddCommand(pipes_listTagsForResourceCmd)
}
