package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Returns all tags for the given Elasticsearch domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_listTagsCmd).Standalone()

		es_listTagsCmd.Flags().String("arn", "", "Specify the `ARN` for the Elasticsearch domain to which the tags are attached that you want to view.")
		es_listTagsCmd.MarkFlagRequired("arn")
	})
	esCmd.AddCommand(es_listTagsCmd)
}
