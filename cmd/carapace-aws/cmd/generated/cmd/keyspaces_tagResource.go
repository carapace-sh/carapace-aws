package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a set of tags with a Amazon Keyspaces resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_tagResourceCmd).Standalone()

	keyspaces_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Keyspaces resource to which to add tags.")
	keyspaces_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon Keyspaces resource.")
	keyspaces_tagResourceCmd.MarkFlagRequired("resource-arn")
	keyspaces_tagResourceCmd.MarkFlagRequired("tags")
	keyspacesCmd.AddCommand(keyspaces_tagResourceCmd)
}
