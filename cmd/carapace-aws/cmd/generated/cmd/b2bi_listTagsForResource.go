package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags associated with the Amazon Resource Name (ARN) that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_listTagsForResourceCmd).Standalone()

		b2bi_listTagsForResourceCmd.Flags().String("resource-arn", "", "Requests the tags associated with a particular Amazon Resource Name (ARN).")
		b2bi_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	b2biCmd.AddCommand(b2bi_listTagsForResourceCmd)
}
