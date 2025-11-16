package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds metadata tags to an elastic cluster resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_tagResourceCmd).Standalone()

		docdbElastic_tagResourceCmd.Flags().String("resource-arn", "", "The ARN identifier of the elastic cluster resource.")
		docdbElastic_tagResourceCmd.Flags().String("tags", "", "The tags that are assigned to the elastic cluster resource.")
		docdbElastic_tagResourceCmd.MarkFlagRequired("resource-arn")
		docdbElastic_tagResourceCmd.MarkFlagRequired("tags")
	})
	docdbElasticCmd.AddCommand(docdbElastic_tagResourceCmd)
}
