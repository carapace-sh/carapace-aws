package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_untagResourceCmd).Standalone()

		ecrPublic_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to delete tags from.")
		ecrPublic_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
		ecrPublic_untagResourceCmd.MarkFlagRequired("resource-arn")
		ecrPublic_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ecrPublicCmd.AddCommand(ecrPublic_untagResourceCmd)
}
