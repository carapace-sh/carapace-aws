package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes the specified set of tags from the specified Elasticsearch domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_removeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_removeTagsCmd).Standalone()

		es_removeTagsCmd.Flags().String("arn", "", "Specifies the `ARN` for the Elasticsearch domain from which you want to delete the specified tags.")
		es_removeTagsCmd.Flags().String("tag-keys", "", "Specifies the `TagKey` list which you want to remove from the Elasticsearch domain.")
		es_removeTagsCmd.MarkFlagRequired("arn")
		es_removeTagsCmd.MarkFlagRequired("tag-keys")
	})
	esCmd.AddCommand(es_removeTagsCmd)
}
