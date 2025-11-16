package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags that are attached to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_untagResourceCmd).Standalone()

		personalize_untagResourceCmd.Flags().String("resource-arn", "", "The resource's Amazon Resource Name (ARN).")
		personalize_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
		personalize_untagResourceCmd.MarkFlagRequired("resource-arn")
		personalize_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	personalizeCmd.AddCommand(personalize_untagResourceCmd)
}
