package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns the tags for an OpenSearch Serverless resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listTagsForResourceCmd).Standalone()

	opensearchserverless_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	opensearchserverless_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	opensearchserverlessCmd.AddCommand(opensearchserverless_listTagsForResourceCmd)
}
