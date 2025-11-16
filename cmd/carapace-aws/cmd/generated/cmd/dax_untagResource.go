package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of tags from a DAX resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_untagResourceCmd).Standalone()

		dax_untagResourceCmd.Flags().String("resource-name", "", "The name of the DAX resource from which the tags should be removed.")
		dax_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys.")
		dax_untagResourceCmd.MarkFlagRequired("resource-name")
		dax_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	daxCmd.AddCommand(dax_untagResourceCmd)
}
