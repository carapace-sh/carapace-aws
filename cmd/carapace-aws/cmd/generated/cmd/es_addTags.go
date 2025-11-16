package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Attaches tags to an existing Elasticsearch domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_addTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_addTagsCmd).Standalone()

		es_addTagsCmd.Flags().String("arn", "", "Specify the `ARN` for which you want to add the tags.")
		es_addTagsCmd.Flags().String("tag-list", "", "List of `Tag` that need to be added for the Elasticsearch domain.")
		es_addTagsCmd.MarkFlagRequired("arn")
		es_addTagsCmd.MarkFlagRequired("tag-list")
	})
	esCmd.AddCommand(es_addTagsCmd)
}
