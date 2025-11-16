package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_tagResourceCmd).Standalone()

	batch_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that tags are added to.")
	batch_tagResourceCmd.Flags().String("tags", "", "The tags that you apply to the resource to help you categorize and organize your resources.")
	batch_tagResourceCmd.MarkFlagRequired("resource-arn")
	batch_tagResourceCmd.MarkFlagRequired("tags")
	batchCmd.AddCommand(batch_tagResourceCmd)
}
