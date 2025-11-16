package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of tags from a Amazon Keyspaces resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_untagResourceCmd).Standalone()

		keyspaces_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Keyspaces resource that the tags will be removed from.")
		keyspaces_untagResourceCmd.Flags().String("tags", "", "A list of existing tags to be removed from the Amazon Keyspaces resource.")
		keyspaces_untagResourceCmd.MarkFlagRequired("resource-arn")
		keyspaces_untagResourceCmd.MarkFlagRequired("tags")
	})
	keyspacesCmd.AddCommand(keyspaces_untagResourceCmd)
}
