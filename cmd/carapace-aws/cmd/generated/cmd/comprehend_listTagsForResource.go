package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with a given Amazon Comprehend resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listTagsForResourceCmd).Standalone()

		comprehend_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the given Amazon Comprehend resource you are querying.")
		comprehend_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	comprehendCmd.AddCommand(comprehend_listTagsForResourceCmd)
}
