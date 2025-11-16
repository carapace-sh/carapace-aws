package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes metadata tags from an elastic cluster resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_untagResourceCmd).Standalone()

		docdbElastic_untagResourceCmd.Flags().String("resource-arn", "", "The ARN identifier of the elastic cluster resource.")
		docdbElastic_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to be removed from the elastic cluster resource.")
		docdbElastic_untagResourceCmd.MarkFlagRequired("resource-arn")
		docdbElastic_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	docdbElasticCmd.AddCommand(docdbElastic_untagResourceCmd)
}
