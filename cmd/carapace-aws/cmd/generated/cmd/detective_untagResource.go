package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_untagResourceCmd).Standalone()

	detective_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the behavior graph to remove the tags from.")
	detective_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys of the tags to remove from the behavior graph.")
	detective_untagResourceCmd.MarkFlagRequired("resource-arn")
	detective_untagResourceCmd.MarkFlagRequired("tag-keys")
	detectiveCmd.AddCommand(detective_untagResourceCmd)
}
