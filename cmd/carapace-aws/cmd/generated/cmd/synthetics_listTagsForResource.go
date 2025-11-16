package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a canary or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_listTagsForResourceCmd).Standalone()

		synthetics_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the canary or group that you want to view tags for.")
		synthetics_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	syntheticsCmd.AddCommand(synthetics_listTagsForResourceCmd)
}
