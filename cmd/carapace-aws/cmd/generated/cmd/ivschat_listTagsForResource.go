package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets information about AWS tags for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_listTagsForResourceCmd).Standalone()

	ivschat_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be retrieved.")
	ivschat_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ivschatCmd.AddCommand(ivschat_listTagsForResourceCmd)
}
