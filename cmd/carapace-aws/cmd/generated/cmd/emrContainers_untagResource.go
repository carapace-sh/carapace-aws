package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_untagResourceCmd).Standalone()

		emrContainers_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of resources.")
		emrContainers_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys of the resources.")
		emrContainers_untagResourceCmd.MarkFlagRequired("resource-arn")
		emrContainers_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	emrContainersCmd.AddCommand(emrContainers_untagResourceCmd)
}
