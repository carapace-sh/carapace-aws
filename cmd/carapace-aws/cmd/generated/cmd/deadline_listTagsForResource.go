package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listTagsForResourceCmd).Standalone()

	deadline_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource ARN to list tags for.")
	deadline_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	deadlineCmd.AddCommand(deadline_listTagsForResourceCmd)
}
