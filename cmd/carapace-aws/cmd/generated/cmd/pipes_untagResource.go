package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified pipes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_untagResourceCmd).Standalone()

	pipes_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the pipe.")
	pipes_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the pipe.")
	pipes_untagResourceCmd.MarkFlagRequired("resource-arn")
	pipes_untagResourceCmd.MarkFlagRequired("tag-keys")
	pipesCmd.AddCommand(pipes_untagResourceCmd)
}
