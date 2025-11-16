package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or set of tags from an OpenSearch Serverless resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_untagResourceCmd).Standalone()

	opensearchserverless_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
	opensearchserverless_untagResourceCmd.Flags().String("tag-keys", "", "The tag or set of tags to remove from the resource.")
	opensearchserverless_untagResourceCmd.MarkFlagRequired("resource-arn")
	opensearchserverless_untagResourceCmd.MarkFlagRequired("tag-keys")
	opensearchserverlessCmd.AddCommand(opensearchserverless_untagResourceCmd)
}
