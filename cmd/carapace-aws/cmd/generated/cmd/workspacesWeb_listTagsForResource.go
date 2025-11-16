package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves a list of tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listTagsForResourceCmd).Standalone()

		workspacesWeb_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		workspacesWeb_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listTagsForResourceCmd)
}
