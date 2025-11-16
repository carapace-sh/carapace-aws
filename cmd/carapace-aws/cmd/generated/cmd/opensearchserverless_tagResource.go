package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates tags with an OpenSearch Serverless resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_tagResourceCmd).Standalone()

		opensearchserverless_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		opensearchserverless_tagResourceCmd.Flags().String("tags", "", "A list of tags (key-value pairs) to add to the resource.")
		opensearchserverless_tagResourceCmd.MarkFlagRequired("resource-arn")
		opensearchserverless_tagResourceCmd.MarkFlagRequired("tags")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_tagResourceCmd)
}
