package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with an Evidently resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_listTagsForResourceCmd).Standalone()

		evidently_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you want to see the tags of.")
		evidently_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	evidentlyCmd.AddCommand(evidently_listTagsForResourceCmd)
}
