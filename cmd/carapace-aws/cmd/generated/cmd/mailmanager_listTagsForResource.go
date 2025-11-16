package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of tags (keys and values) assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listTagsForResourceCmd).Standalone()

		mailmanager_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to retrieve tags from.")
		mailmanager_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mailmanagerCmd.AddCommand(mailmanager_listTagsForResourceCmd)
}
