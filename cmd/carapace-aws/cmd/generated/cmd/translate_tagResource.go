package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a specific tag with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_tagResourceCmd).Standalone()

		translate_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the given Amazon Translate resource to which you want to associate the tags.")
		translate_tagResourceCmd.Flags().String("tags", "", "Tags being associated with a specific Amazon Translate resource.")
		translate_tagResourceCmd.MarkFlagRequired("resource-arn")
		translate_tagResourceCmd.MarkFlagRequired("tags")
	})
	translateCmd.AddCommand(translate_tagResourceCmd)
}
