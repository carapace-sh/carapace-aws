package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_untagResourceCmd).Standalone()

		codestarConnections_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
		codestarConnections_untagResourceCmd.Flags().String("tag-keys", "", "The list of keys for the tags to be removed from the resource.")
		codestarConnections_untagResourceCmd.MarkFlagRequired("resource-arn")
		codestarConnections_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_untagResourceCmd)
}
