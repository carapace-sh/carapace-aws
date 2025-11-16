package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from an App Runner resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_untagResourceCmd).Standalone()

		apprunner_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove tags from.")
		apprunner_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys that you want to remove.")
		apprunner_untagResourceCmd.MarkFlagRequired("resource-arn")
		apprunner_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	apprunnerCmd.AddCommand(apprunner_untagResourceCmd)
}
