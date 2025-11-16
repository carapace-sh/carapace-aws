package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_tagResourceCmd).Standalone()

	pipes_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the pipe.")
	pipes_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs associated with the pipe.")
	pipes_tagResourceCmd.MarkFlagRequired("resource-arn")
	pipes_tagResourceCmd.MarkFlagRequired("tags")
	pipesCmd.AddCommand(pipes_tagResourceCmd)
}
