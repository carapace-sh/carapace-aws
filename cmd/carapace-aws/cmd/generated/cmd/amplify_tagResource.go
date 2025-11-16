package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags the resource with a tag key and value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_tagResourceCmd).Standalone()

		amplify_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to use to tag a resource.")
		amplify_tagResourceCmd.Flags().String("tags", "", "The tags used to tag the resource.")
		amplify_tagResourceCmd.MarkFlagRequired("resource-arn")
		amplify_tagResourceCmd.MarkFlagRequired("tags")
	})
	amplifyCmd.AddCommand(amplify_tagResourceCmd)
}
