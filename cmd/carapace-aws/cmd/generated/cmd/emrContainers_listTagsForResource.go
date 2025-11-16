package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to the resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_listTagsForResourceCmd).Standalone()

		emrContainers_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of tagged resources.")
		emrContainers_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	emrContainersCmd.AddCommand(emrContainers_listTagsForResourceCmd)
}
