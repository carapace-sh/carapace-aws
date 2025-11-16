package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for the AWS resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_tagResourceCmd).Standalone()

		ivschat_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be tagged.")
		ivschat_tagResourceCmd.Flags().String("tags", "", "Array of tags to be added or updated.")
		ivschat_tagResourceCmd.MarkFlagRequired("resource-arn")
		ivschat_tagResourceCmd.MarkFlagRequired("tags")
	})
	ivschatCmd.AddCommand(ivschat_tagResourceCmd)
}
