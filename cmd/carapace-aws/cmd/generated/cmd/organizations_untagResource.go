package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes any tags with the specified keys from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_untagResourceCmd).Standalone()

		organizations_untagResourceCmd.Flags().String("resource-id", "", "The ID of the resource to remove a tag from.")
		organizations_untagResourceCmd.Flags().String("tag-keys", "", "The list of keys for tags to remove from the specified resource.")
		organizations_untagResourceCmd.MarkFlagRequired("resource-id")
		organizations_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	organizationsCmd.AddCommand(organizations_untagResourceCmd)
}
