package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add tags to, or update the tag values of, an App Runner resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_tagResourceCmd).Standalone()

		apprunner_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to update tags for.")
		apprunner_tagResourceCmd.Flags().String("tags", "", "A list of tag key-value pairs to add or update.")
		apprunner_tagResourceCmd.MarkFlagRequired("resource-arn")
		apprunner_tagResourceCmd.MarkFlagRequired("tags")
	})
	apprunnerCmd.AddCommand(apprunner_tagResourceCmd)
}
