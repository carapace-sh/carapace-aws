package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a specific tag with an Amazon Comprehend resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_tagResourceCmd).Standalone()

	comprehend_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the given Amazon Comprehend resource to which you want to associate the tags.")
	comprehend_tagResourceCmd.Flags().String("tags", "", "Tags being associated with a specific Amazon Comprehend resource.")
	comprehend_tagResourceCmd.MarkFlagRequired("resource-arn")
	comprehend_tagResourceCmd.MarkFlagRequired("tags")
	comprehendCmd.AddCommand(comprehend_tagResourceCmd)
}
