package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes metadata tags from a FinSpace resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_untagResourceCmd).Standalone()

		finspace_untagResourceCmd.Flags().String("resource-arn", "", "A FinSpace resource from which you want to remove a tag or tags.")
		finspace_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys (names) of one or more tags to be removed.")
		finspace_untagResourceCmd.MarkFlagRequired("resource-arn")
		finspace_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	finspaceCmd.AddCommand(finspace_untagResourceCmd)
}
