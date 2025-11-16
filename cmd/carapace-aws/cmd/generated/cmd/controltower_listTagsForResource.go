package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags associated with the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listTagsForResourceCmd).Standalone()

	controltower_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	controltower_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	controltowerCmd.AddCommand(controltower_listTagsForResourceCmd)
}
