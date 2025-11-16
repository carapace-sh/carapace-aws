package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on a elastic cluster resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_listTagsForResourceCmd).Standalone()

		docdbElastic_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN identifier of the elastic cluster resource.")
		docdbElastic_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	docdbElasticCmd.AddCommand(docdbElastic_listTagsForResourceCmd)
}
