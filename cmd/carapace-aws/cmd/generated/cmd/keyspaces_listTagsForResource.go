package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of all tags associated with the specified Amazon Keyspaces resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_listTagsForResourceCmd).Standalone()

		keyspaces_listTagsForResourceCmd.Flags().String("max-results", "", "The total number of tags to return in the output.")
		keyspaces_listTagsForResourceCmd.Flags().String("next-token", "", "The pagination token.")
		keyspaces_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Keyspaces resource.")
		keyspaces_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	keyspacesCmd.AddCommand(keyspaces_listTagsForResourceCmd)
}
