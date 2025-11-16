package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets information about Amazon Web Services tags for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listTagsForResourceCmd).Standalone()

	ivs_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be retrieved.")
	ivs_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ivsCmd.AddCommand(ivs_listTagsForResourceCmd)
}
