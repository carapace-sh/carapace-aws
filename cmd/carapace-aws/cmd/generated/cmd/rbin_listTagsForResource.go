package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to a retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_listTagsForResourceCmd).Standalone()

	rbin_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the retention rule.")
	rbin_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	rbinCmd.AddCommand(rbin_listTagsForResourceCmd)
}
